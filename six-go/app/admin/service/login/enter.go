package loginService

import (
	"errors"
	"time"

	"six-go/database/models"
)

func New(token string, keepalive ...bool) Driver {
	var loginStorage RedisDriver
	loginStorage.token = token
	if len(keepalive) > 0 && keepalive[0] {
		// 调用保活函数
		loginStorage.keepalive()
	}
	return loginStorage
}

func ClearAll(username string) {
	var loginStorage RedisDriver
	loginStorage.username = username
	loginStorage.ClearAll()
}

// Driver 任何存储登录状态的驱动只要实现该接口 即可无缝切换 只需要将 LoginStorage 函数中变量 loginStorage 的类型 修改为实现 该接口的 Struct 即可
type Driver interface {
	IsLogin() (isLogin bool)
	Get() (userinfo StorageValue)
	Set(userinfo *models.AuthUser) error
	keepalive()
	Clear()
	ClearAll()
}

type sessionConfig struct {
	// userLimit 设备登录限制数 userLimit = 1 单设备登录
	userLimit int

	// userLimitOverMode 登录人数溢出模式
	// 阻塞模式：userLimitOverModeBlock  会阻止第 userLimit + 1 个用户登录
	// 覆盖模式: userLimitOverModeCover 会覆盖最早过期的用户
	userLimitOverMode int

	// expireTime 过期时间
	expireTime time.Duration
}

type StorageValue struct {
	expire   int64
	IsLogin  bool
	Userinfo *models.AuthUser
}

const (
	userLimitOverModeBlock = 0
	userLimitOverModeCover = 1
)

var (
	ErrorLoginBlocked = errors.New("该账号登录设备超出系统限制，请在其他设备退出登录后重试。")
)

var loginConfig = sessionConfig{
	userLimit:         1,
	userLimitOverMode: userLimitOverModeCover,
	expireTime:        time.Hour * 24 * 7, // 默认每个token登录有效期保活7天
}
