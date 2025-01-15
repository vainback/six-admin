package db

import "gorm.io/gorm"

type ModelIof interface {
	TableName() string                            // 表名
	KeywordFields() []string                      // 模糊查询字段
	FilterSqlBuilder() func(db *gorm.DB) *gorm.DB // 过滤查询
	GetId() int64
}
