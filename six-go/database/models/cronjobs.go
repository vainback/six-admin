package models

import (
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"six-go/database/db"
)

var TableCronJob CronJob

type CronJob struct {
	db.MODEL
	db.SoftDelete
	Name   string `json:"name"`  // 唯一标识
	Title  string `json:"title"` // 标题
	Times  string `json:"times"` // 定时表达式
	Status int    `json:"status" gorm:"type:tinyint(1);comment:状态 -1 禁用 1启用"`
}

func (data CronJob) TableName() string {
	return "cron_jobs"
}

func (data CronJob) HasTableName(filed string) string {
	return strings.Join([]string{data.TableName(), filed}, ".")
}

func (data CronJob) KeywordFields() []string {
	return []string{
		data.HasTableName("name"),
		data.HasTableName("title"),
	}
}

func (data CronJob) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.Name != "" {
			db = db.Where(data.HasTableName("type")+" = ?", data.Name)
		}
		if data.Title != "" {
			db = db.Where(data.HasTableName("label")+" LIKE ?", "%"+data.Title+"%")
		}

		if data.Status != 0 {
			db = db.Where("status = ?", data.Status)
		}
		return db
	}
}

func (data CronJob) GetId() int64 {
	return data.Id
}

func (data CronJob) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return validation.ValidateStruct(&data,
		validation.Field(&data.Name, validation.Required.Error("唯一标识不能为空")),
		validation.Field(&data.Title, validation.Required.Error("标题不能为空")),
		validation.Field(&data.Times, validation.Required.Error("定时时间不能为空")),
	)
}
