package models

import (
	"errors"
	"github.com/vainback/six-util/v3"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"six-go/database/db"
)

const RootRoleId = -20240929

var TableAuthRole AuthRole

// AuthRole 角色
type AuthRole struct {
	db.MODEL
	db.SoftDelete
	TenantIdentify
	ParentId int64  `json:"parent_id" gorm:"comment:父级id"`
	Title    string `json:"title" gorm:"comment:角色名"`
	Sign     string `json:"sign" gorm:"not null;unique;comment:角色签名"`
	Status   int    `json:"status" gorm:"type:tinyint(1);comment:状态 -1 禁用 1启用"`
}

func (data AuthRole) TableName() string {
	return "auth_role"
}

func (data AuthRole) HasTableName(filed string) string {
	return six.Strings(data.TableName(), ".", filed).String()
}

func (data AuthRole) KeywordFields() []string {
	return six.Arrays(
		data.HasTableName("title"),
		data.HasTableName("sign"),
	)
}

func (data AuthRole) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.ParentId != 0 {
			db = db.Where(data.HasTableName("parent_id")+" = ?", data.ParentId)
		}
		if data.Title != "" {
			db = db.Where(data.HasTableName("title")+" LIKE ?", "%"+data.Title+"%")
		}
		if data.Sign != "" {
			db = db.Where(data.HasTableName("sign")+" LIKE ?", "%"+data.Sign+"%")
		}
		if data.Status != 0 {
			db = db.Where(data.HasTableName("status")+" = ?", data.Status)
		}
		return db.Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))
	}
}

func (data AuthRole) GetId() int64 {
	return data.Id
}

func (data AuthRole) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return validation.ValidateStruct(&data,
		validation.Field(&data.ParentId, validation.Min(0).Error("上级菜单不合法")),
		validation.Field(&data.Title, validation.Required.Error("角色名称必填")),
		validation.Field(&data.Sign, validation.Required.Error("角色签名必填")),
		validation.Field(&data.Status, validation.Required.Error("角色状态必选")),
	)
}

func (data AuthRole) HasChildren() (childNums int64, err error) {
	err = db.DB().Model(&AuthRole{}).Where("parent_id = ?", data.Id).Count(&childNums).Error
	return
}
