package models

import (
	"strings"

	"gorm.io/gorm"
)

type TenantIdentify struct {
	TenantId int64 `json:"tenant_id" gorm:"column:tenant_id;comment:租户id;index"`
}

func (data TenantIdentify) FiledName() string {
	return "tenant_id"
}

// SqlBuilder tables 只取第一个参数
func (data TenantIdentify) SqlBuilder(tableName ...string) func(db *gorm.DB) *gorm.DB {
	var tableQuery strings.Builder
	if len(tableName) > 0 {
		tableQuery.WriteString(tableName[0])
		tableQuery.WriteString(".")
	}
	tableQuery.WriteString("tenant_id = ?")
	return func(db *gorm.DB) *gorm.DB {
		if data.TenantId > 0 {
			db = db.Where(tableQuery.String(), data.TenantId)
		}
		return db
	}
}
