package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
	"six-go/database/db"
)

var TableFiles Files

type Files struct {
	db.MODEL
	db.SoftDelete
	TenantIdentify
	Mime string `json:"mime" gorm:"index"` // 1 图片 2 视频 3 音频 4 其他文件
	Type string `json:"type" gorm:"index"` // 存储方式 local | remote  本地|远端
	Url  string `json:"url"`               // 文件url
}

func (data Files) TableName() string {
	return "files"
}

func (data Files) HasTableName(filed string) string {
	return strings.Join([]string{data.TableName(), filed}, ".")
}

func (data Files) KeywordFields() []string {
	return []string{}
}

func (data Files) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}

		if data.Mime != "" {
			return db.Where(data.HasTableName("id")+" = ?", data.Mime)
		}

		if data.Type != "" {
			return db.Where(data.HasTableName("id")+" = ?", data.Type)
		}

		return db.Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))
	}
}

func (data Files) GetId() int64 {
	return data.Id
}

func (data Files) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}
	return nil
}
