package models

import (
	"errors"
	"github.com/vainback/six-util/v3"
	"six-go/utils"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"six-go/database/db"
	"six-go/database/v_errors"
)

var TableAuthUser AuthUser

type AuthUser struct {
	db.MODEL
	db.SoftDelete
	TenantIdentify                // 租户
	Username       string         `json:"username" gorm:"comment:用户名 可以以手机号作为用户名;unique"`
	Password       utils.Password `json:"password" gorm:"comment:密码"`
	Nickname       string         `json:"nickname" gorm:"comment:昵称"`
	Status         int            `json:"status" gorm:"type:tinyint(1);comment:状态 -1 禁用 1启用"`
	RoleIds        []int64        `json:"role_ids" gorm:"serializer:json;comment:角色ids"`
	IsRoot         bool           `json:"is_root" gorm:"<-:false;default:0;comment:是否是root用户 1 是 0 否"`
	OrganizationId int64          `json:"organization_id" gorm:"index"` // 组织部门Id
	JobId          int64          `json:"job_id" gorm:"index"`          // 职位Id
	// 其他非必要字段 需要时再进行扩展
}

func (data AuthUser) TableName() string {
	return "auth_user"
}

func (data AuthUser) HasTableName(filed string) string {
	return six.Strings(data.TableName(), ".", filed).String()
}

func (data AuthUser) KeywordFields() []string {
	return six.Arrays(
		data.HasTableName("username"),
		data.HasTableName("nickname"),
	)
}

func (data AuthUser) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.Username != "" {
			db = db.Where(data.HasTableName("username")+" = ?", data.Username)
		}
		if data.Nickname != "" {
			db = db.Where(data.HasTableName("nickname")+" LIKE ?", "%"+data.Nickname+"%")
		}
		if data.Status != 0 {
			db = db.Where(data.HasTableName("status")+" = ?", data.Status)
		}

		if data.JobId != 0 {
			db = db.Where(data.HasTableName("job_id")+" = ?", data.JobId)
		}

		if data.OrganizationId != 0 {
			// sons := NewGetOrganizationSons(data.OrganizationId)
			sons := NewSon(data.OrganizationId).GetIds(TableOrganization)
			db = db.Where(data.HasTableName("organization_id")+" in ?", append(sons, data.OrganizationId))
		}

		return db.Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))
	}
}

func (data AuthUser) GetId() int64 {
	return data.Id
}

func (data AuthUser) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return validation.ValidateStruct(&data,
		validation.Field(&data.Username,
			validation.Required.Error("用户名不能为空"),
			validation.By(func(value interface{}) error {
				var count int64
				if err := db.DB().Model(&AuthUser{}).Where("id <> ?", data.Id).Where("username=?", data.Username).Count(&count).Error; err != nil {
					return err
				}
				if count > 0 {
					return v_errors.Unique("用户名")
				}
				return nil
			})),
		validation.Field(&data.Password,
			// validation.Required.Error("密码不能为空"),
			// 当data.Id为0时 密码必填
			validation.By(func(value interface{}) error {
				if data.Id == 0 {
					if validation.Required.Validate(data.Password) != nil {
						return errors.New("密码不能为空")
					}
				}
				return nil
			}),
			validation.Length(6, 16).Error("密码长度需在6-16位之间"),
		),
		validation.Field(&data.Nickname, validation.Required.Error("昵称不能为空")),
		validation.Field(&data.Status, validation.Required.Error("菜单状态必选")),
		validation.Field(&data.RoleIds, validation.Required.Error("角色必选")),
	)
}
