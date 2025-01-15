package loginService

import (
	"sync"
	"time"

	"six-go/database/models"
)

type LocalDriver struct {
	token    string
	username string
}

var localstorage sync.Map

func (ts LocalDriver) IsLogin() bool {
	return ts.Get().IsLogin
}

func (ts LocalDriver) Get() (storageValue StorageValue) {
	value, ok := localstorage.Load(ts.token)
	if !ok {
		return
	}
	return value.(StorageValue)
}

func (ts LocalDriver) Set(userinfo *models.AuthUser) error {
	var alreadyTokens []string
	value, ok := localstorage.Load(userinfo.Username)
	if ok && value != nil {
		alreadyTokens = value.([]string)
	}

	if len(alreadyTokens) >= loginConfig.userLimit {
		switch loginConfig.userLimitOverMode {
		case userLimitOverModeBlock: // 阻塞模式
			return ErrorLoginBlocked
		default:
			// 默认为覆盖模式
			// 移除最早过期的用户
			var expire int64
			var tsToken string
			for index, token := range alreadyTokens {
				tsExpire := (LocalDriver{token: token}).Get().expire
				if index == 0 || expire > tsExpire {
					expire = tsExpire
					tsToken = token
				}
			}
			(LocalDriver{token: tsToken, username: userinfo.Username}).Clear()
			alreadyTokens = append(alreadyTokens, ts.token)
		}
	}

	localstorage.Store(ts.token, StorageValue{expire: time.Now().Add(loginConfig.expireTime).Unix(), Userinfo: userinfo, IsLogin: true})
	localstorage.Store(userinfo.Username, alreadyTokens)
	return nil
}

func (ts LocalDriver) keepalive() {
	store := ts.Get()
	if store.IsLogin {
		if store.expire > time.Now().Unix() {
			store.expire = time.Now().Add(loginConfig.expireTime).Unix()
			localstorage.Store(ts.token, store)
		} else {
			// 执行清除
			ts.username = store.Userinfo.Username
			ts.Clear()
		}
	}
}

func (ts LocalDriver) Clear() {
	localstorage.Delete(ts.token)

	var alreadyTokens []string
	value, ok := localstorage.Load(ts.username)
	if ok && value != nil {
		alreadyTokens = value.([]string)
	}

	if len(alreadyTokens) > 0 {
		var newTokens = make([]string, len(alreadyTokens)-1)
		for _, token := range alreadyTokens {
			if token != ts.token {
				newTokens = append(newTokens, token)
			}
		}
		localstorage.Store(ts.username, newTokens)
	}

	ts.username = ""
	ts.token = ""
}

func (ts LocalDriver) ClearAll() {
	localstorage.Delete(ts.username)
}
