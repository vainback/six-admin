package models

import (
	"errors"
	"gorm.io/gorm"
	"six-go/database/db"
)

var TableCodeGenerator CodeGenerator

type CodeGenerator struct {
	db.MODEL
	db.SoftDelete
	Table        string                 `json:"table" gorm:"column:table"`
	Fields       *[]CodeGeneratorFields `json:"fields" gorm:"column:fields;serializer:json;type:longtext"`
	Title        string                 `json:"title" gorm:"column:title"`                 // 菜单名称
	ParentModule int64                  `json:"parent_module" gorm:"column:parent_module"` // 父级菜单id
	IsSoftDelete bool                   `json:"is_soft_delete" gorm:"column:is_soft_delete"`
	IsTenant     bool                   `json:"is_tenant" gorm:"column:is_tenant"`
}

type CodeGeneratorFields struct {
	Title      string  `json:"title"`   // 展示名
	Name       string  `json:"name"`    // 字段名
	Type       string  `json:"type"`    // 字段类型 int int64 uint uint64 string float64 decimal select radio checkbox switch textarea 上传图片 多图上传 上传文件 时间选择器 日期选择器 日期时间选择器 富文本
	Default    *string `json:"default"` // 字段默认值 nil 为 null 非nil is not null
	Index      string  `json:"index"`   // 索引 可选值为 unique | index
	Comment    string  `json:"comment"` // 注释
	IsKeyword  bool    `json:"is_keyword"`
	IsRequired bool    `json:"is_required"`
	DictType   string  `json:"dict_type"` // 字典类型
}

func (data CodeGenerator) KeywordFields() []string {
	return []string{"`table`", "title"}
}

func (data CodeGenerator) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func (data CodeGenerator) GetId() int64 {
	return data.Id
}

func (data CodeGenerator) TableName() string {
	return "code_generator"
}

func (data CodeGenerator) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return nil
}
