package loginService

import (
	"time"

	"six-go/database/db"
	"six-go/database/models"
)

type MysqlDriver struct {
	token    string
	username string
}

func (ts MysqlDriver) IsLogin() (isLogin bool) {
	isLogin = ts.Get().IsLogin
	return
}

func (ts MysqlDriver) Get() (storageValue StorageValue) {
	var value models.AuthUserLoginStorage
	if err := db.DB().First(&value, "token = ?", ts.token).Error; err != nil {
		return
	}

	var user models.AuthUser
	if err := db.DB().First(&user, "username = ?", value.Username).Error; err != nil {
		return
	}

	return StorageValue{
		IsLogin:  true,
		Userinfo: &user,
	}
}

func (ts MysqlDriver) Set(userinfo *models.AuthUser) error {
	var alreadyTokens []models.AuthUserLoginStorage
	if err := db.DB().Where("token = ?", ts.token).Order("expire asc").Find(&alreadyTokens).Error; err != nil {
		return err
	}

	if len(alreadyTokens) >= loginConfig.userLimit {
		switch loginConfig.userLimitOverMode {
		case userLimitOverModeBlock: // 阻塞模式
			return ErrorLoginBlocked
		default: // 默认为覆盖模式 移除最早过期的用户
			(MysqlDriver{token: alreadyTokens[0].Token}).Clear()
		}
	}

	return db.DB().Create(&models.AuthUserLoginStorage{
		Username: userinfo.Username,
		Token:    ts.token,
		Expire:   time.Now().Add(loginConfig.expireTime).Unix(),
	}).Error
}

func (ts MysqlDriver) keepalive() {
	// 清除所有已失效的token
	_ = db.DB().Delete(&models.AuthUserLoginStorage{}, "expire <= ?", time.Now().Unix()).Error
	// 保活当前token
	_ = db.DB().Model(&models.AuthUserLoginStorage{}).Where("token = ?", ts.token).Update("expire", time.Now().Add(loginConfig.expireTime).Unix()).Error
}

// Clear 清除指定token
func (ts MysqlDriver) Clear() {
	_ = db.DB().Delete(&models.AuthUserLoginStorage{Token: ts.token}).Error
}

func (ts MysqlDriver) ClearAll() {
	_ = db.DB().Delete(&models.AuthUserLoginStorage{Username: ts.username}).Error
}
