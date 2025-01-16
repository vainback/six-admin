package entity

import (
	"fmt"
	"github.com/vainback/six-util/v3"
	"strings"

	"gorm.io/gorm"
	"six-go/database/models"
)

type LogJoinUser struct {
	models.OperateLog
	User models.AuthUser `json:"user" gorm:"embedded;embeddedPrefix:user_"`
}

func (e LogJoinUser) TableName() string {
	return e.OperateLog.TableName()
}

func (e LogJoinUser) KeywordFields() []string {
	return append(e.OperateLog.KeywordFields(), e.User.KeywordFields()...)
}

func (e LogJoinUser) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(e.OperateLog.FilterSqlBuilder()).Scopes(e.User.FilterSqlBuilder())
	}
}

func (e LogJoinUser) Fields() string {
	var fields = []string{
		six.Strings(e.OperateLog.TableName(), ".*").String(),
		six.Strings(e.User.TableName(), ".username as user_username").String(),
		six.Strings(e.User.TableName(), ".nickname as user_nickname").String(),
	}
	return strings.Join(fields, ",")
}

func (e LogJoinUser) LeftJoin() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(e.OperateLog.TableName()).
			Joins(fmt.Sprintf("LEFT JOIN %s ON %s.id = %s.uid", e.User.TableName(), e.User.TableName(), e.OperateLog.TableName()))
	}
}

func (e LogJoinUser) Select() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(e.Fields())
	}
}
