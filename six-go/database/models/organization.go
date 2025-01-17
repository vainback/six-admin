package models

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vainback/six-util/v3"
	"gorm.io/gorm"
	"six-go/database/db"
)

var TableOrganization Organization

type Organization struct {
	db.MODEL
	db.SoftDelete
	TenantIdentify
	ParentId    int64  `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (data Organization) TableName() string {
	return "organization"
}

func (data Organization) HasTableName(filed string) string {
	return six.Strings(data.TableName(), ".", filed).String()
}

func (data Organization) KeywordFields() []string {
	return six.Arrays(
		data.HasTableName("name"),
		data.HasTableName("description"),
	)
}

func (data Organization) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.ParentId != 0 {
			db = db.Where(data.HasTableName("parent_id")+" = ?", data.ParentId)
		}
		if data.Name != "" {
			db = db.Where(data.HasTableName("title")+" LIKE ?", "%"+data.Name+"%")
		}
		if data.Description != "" {
			db = db.Where(data.HasTableName("sign")+" LIKE ?", "%"+data.Description+"%")
		}

		return db.Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))
	}
}

func (data Organization) GetId() int64 {
	return data.Id
}

func (data Organization) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return validation.ValidateStruct(&data,
		validation.Field(&data.ParentId, validation.Min(0).Error("上级组织不合法")),
		validation.Field(&data.Name, validation.Required.Error("角色名称必填")),
	)
}

func (data Organization) HasChildren() (childNums int64, err error) {
	err = db.DB().Model(&Organization{}).Where("parent_id = ?", data.Id).Count(&childNums).Error
	return
}
