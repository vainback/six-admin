package models

import (
	"errors"
	"github.com/vainback/six-util/v3"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
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
	err = db.DB().Model(&Organization{}).Where("id = ?", data.Id).Count(&childNums).Error
	return
}

//
// type GetOrganizationSons struct {
// 	list []Organization
// }
//
// func NewGetOrganizationSons(id int64) []int64 {
// 	var list []Organization
// 	if err := db.DB().Find(&list).Error; err != nil {
// 		log.Println(err)
// 		return nil
// 	}
//
// 	var g = &GetOrganizationSons{list: list}
// 	return g.parserSons(id)
// }
//
// func (g *GetOrganizationSons) parserSons(id int64) []int64 {
// 	if len(g.list) == 0 {
// 		return nil
// 	}
// 	var sons []int64
// 	for _, v := range g.list {
// 		if v.ParentId == id {
// 			if childs := g.parserSons(v.Id); len(childs) > 0 {
// 				sons = append(sons, childs...)
// 			}
// 			sons = append(sons, v.Id)
// 		}
// 	}
// 	return sons
// }
