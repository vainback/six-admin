package models

import (
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"six-go/database/db"
)

var TableTenant Tenant

type Tenant struct {
	db.MODEL
	db.SoftDelete
	Name   string `json:"name" gorm:"comment:租户名"`
	Sign   string `json:"sign" gorm:"unique;comment:租户唯一标识"`
	Domain string `json:"domain" gorm:"comment:单独的域名"`
	Status int    `json:"status" gorm:"type:tinyint(1);comment:状态 -1 禁用 1启用"`
}

func (data Tenant) TableName() string {
	return "tenant"
}

func (data Tenant) HasTableName(filed string) string {
	return strings.Join([]string{data.TableName(), filed}, ".")
}

func (data Tenant) KeywordFields() []string {
	return []string{
		data.HasTableName("name"),
		data.HasTableName("sign"),
		data.HasTableName("domain"),
	}
}

func (data Tenant) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.Name != "" {
			db = db.Where(data.HasTableName("name")+" like ?", "%"+data.Name+"%")
		}
		if data.Sign != "" {
			db = db.Where(data.HasTableName("sign")+" like ?", "%"+data.Sign+"%")
		}
		if data.Domain != "" {
			db = db.Where(data.HasTableName("domain")+" like ?", "%"+data.Domain+"%")
		}

		if data.Status != 0 {
			db = db.Where(data.HasTableName("status")+" = ?", data.Status)
		}
		return db
	}
}

func (data Tenant) GetId() int64 {
	return data.Id
}

func (data Tenant) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return validation.ValidateStruct(&data,
		validation.Field(&data.Name, validation.Required.Error("租户名必须填写")),
		validation.Field(&data.Sign, validation.Required.Error("租户Sign必须填写")),
	)
}
