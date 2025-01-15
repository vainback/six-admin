package loginService

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"six-go/database/models"
	"six-go/extra/xredis"
	"six-go/utils/arrays"
)

type RedisDriver struct {
	token    string
	username string
	oldToken string
}

func (ts RedisDriver) IsLogin() bool {
	return xredis.IsExpire(ts.realTokenKey(), 5)
}

func (ts RedisDriver) Get() (userinfo StorageValue) {
	value := xredis.GetString(ts.realTokenKey())
	if value == "" {
		return
	}

	if err := jsoniter.UnmarshalFromString(value, &userinfo); err != nil {
		log.Println(err)
	}
	return
}

func (ts RedisDriver) Set(userinfo *models.AuthUser) error {
	ts.username = userinfo.Username
	oldToken, tokens := ts.clear()
	if len(tokens) >= loginConfig.userLimit {
		switch loginConfig.userLimitOverMode {
		case userLimitOverModeBlock: // 阻塞模式
			return ErrorLoginBlocked
		default:
			tokens = arrays.NewOrdered(tokens...).Replace(oldToken, ts.token)
			if err := xredis.Del(redisRealTokenKey(oldToken)); err != nil {
				log.Println(err)
			}
		}
	} else {
		tokens = append(tokens, ts.token)
	}

	setValue, _ := jsoniter.Marshal(StorageValue{Userinfo: userinfo, IsLogin: true})
	if err := xredis.SetExp(ts.realTokenKey(), string(setValue), loginConfig.expireTime); err != nil {
		return err
	}

	tokensBytes, _ := jsoniter.Marshal(tokens)
	return xredis.Set(ts.realUsernameKey(), tokensBytes)
}

func (ts RedisDriver) keepalive() {
	if ts.IsLogin() {
		_ = xredis.Expire(ts.realTokenKey(), loginConfig.expireTime)
	}
}

func (ts RedisDriver) Clear() {
	_ = xredis.Del(ts.realTokenKey())
}

func (ts RedisDriver) ClearAll() {
	var alreadyTokens []string
	value := xredis.GetString(ts.realUsernameKey())
	if value != "" {
		if err := jsoniter.UnmarshalFromString(value, &alreadyTokens); err != nil {
			return
		}
	}
	if len(alreadyTokens) > 0 {
		for _, token := range alreadyTokens {
			_ = xredis.Del(redisRealTokenKey(token))
		}
	}
	_ = xredis.Del(ts.realUsernameKey())
}

func (ts RedisDriver) clear() (oldToken string, tokens []string) {
	var alreadyTokens []string
	value := xredis.GetString(ts.realUsernameKey())
	if value != "" {
		if err := jsoniter.UnmarshalFromString(value, &alreadyTokens); err != nil {
			return
		}
	}
	tokens = make([]string, 0)
	if len(alreadyTokens) > 0 {
		var oldTokenSeconds = float64(7 * 24 * 60 * 60)
		for _, token := range alreadyTokens {
			seconds := xredis.HasExpireSeconds(redisRealTokenKey(token))
			if seconds == -1 || seconds > 5 {
				tokens = append(tokens, token)
			} else {
				_ = xredis.Del(ts.realTokenKey())
			}

			if seconds < oldTokenSeconds {
				oldTokenSeconds = seconds
				oldToken = token
			}
		}
	}
	_ = xredis.Set(ts.realUsernameKey(), tokens)
	return
}

func redisRealTokenKey(t string) string {
	return fmt.Sprintf("%s:admin-user-token:%s", xredis.KeyPrefix, t)
}

func realUsernameKey(username string) string {
	return fmt.Sprintf("%s:admin-username:%s", xredis.KeyPrefix, username)
}

func (ts RedisDriver) realTokenKey() string {
	return fmt.Sprintf("%s:admin-user-token:%s", xredis.KeyPrefix, ts.token)
}

func (ts RedisDriver) realUsernameKey() string {
	return fmt.Sprintf("%s:admin-username:%s", xredis.KeyPrefix, ts.username)
}
