package models

import (
	"six-go/database/db"
)

var TableAuthUserLoginStorage AuthUserLoginStorage

type AuthUserLoginStorage struct {
	db.MODEL
	Username string `json:"username" gorm:"comment:username"`
	Token    string `json:"token" gorm:"comment:token"`
	Expire   int64  `json:"expire" gorm:"column:expire"`
}

func (data AuthUserLoginStorage) TableName() string {
	return "auth_user_login_storage"
}
