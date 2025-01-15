package models

import (
	"strings"

	"gorm.io/gorm"
	"six-go/database/db"
)

var TableAuthRelation AuthRelation

// AuthRelation 角色 - 菜单规则 关联表
type AuthRelation struct {
	db.MODEL
	TenantIdentify
	RoleId int64 `json:"role_id"`
	RuleId int64 `json:"rule_id"`
}

func (data AuthRelation) TableName() string {
	return "auth_relation"
}

func (data AuthRelation) HasTableName(filed string) string {
	return strings.Join([]string{data.TableName(), filed}, ".")
}

func (data AuthRelation) KeywordFields() []string {
	return []string{}
}

func (data AuthRelation) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.RoleId != 0 {
			db = db.Where(data.HasTableName("role_id")+" = ?", data.RoleId)
		}
		if data.RuleId != 0 {
			db = db.Where(data.HasTableName("rule_id")+" = ?", data.RuleId)
		}

		return db.Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))
	}
}

func (data AuthRelation) GetId() int64 {
	return data.Id
}
