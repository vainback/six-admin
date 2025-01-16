package models

import (
	"errors"
	"github.com/vainback/six-util/v3"
	"log"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"six-go/database/db"
)

var TableAuthRule AuthRule

type AuthRule struct {
	db.MODEL
	db.SoftDelete
	ParentId  int64     `json:"parent_id" gorm:"not null;default:0;comment:父级ID"`
	Type      RuleTypes `json:"type" gorm:"comment:菜单类型"`
	Component string    `json:"component" gorm:"not null;comment:组件路径"`
	Name      string    `json:"name" gorm:"comment:路由name"`
	Path      string    `json:"path" gorm:"comment:路由path"`
	Auth      string    `json:"auth" gorm:"comment:权限标识"`
	ApiRoute  string    `json:"api_route" gorm:"comment:API路由路径"`
	Locale    string    `json:"locale" gorm:"column:locale;comment:菜单名称"` // 菜单名称
	Icon      string    `json:"icon" gorm:"comment:菜单图标"`                 //
	Order     int       `json:"order" gorm:"column:order;comment:显示排序 数字越小越靠前"`
	Status    int       `json:"status" gorm:"type:tinyint(4);comment:菜单状态 -1 禁用 1 正常"`
}

type AuthRuleForMenu struct {
	AuthRule
	Roles []string `json:"roles"`
}

type RuleTypes string

const (
	RuleTypesMenu   RuleTypes = "menu"   // 无url菜单
	RuleTypesPage   RuleTypes = "page"   // url菜单
	RuleTypesButton RuleTypes = "button" // 按钮
)

func (AuthRule) TableName() string {
	return "auth_rule"
}

func (data AuthRule) HasTableName(filed string) string {
	return six.Strings(data.TableName(), ".", filed).String()
}

func (data AuthRule) KeywordFields() []string {
	return []string{
		data.HasTableName("name"),
		data.HasTableName("path"),
		data.HasTableName("locale"),
	}
}

func (data AuthRule) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.ParentId != 0 {
			db = db.Where(data.HasTableName("parent_id")+" = ?", data.ParentId)
		}
		if data.Type != "" {
			db = db.Where(data.HasTableName("type")+" = ?", data.Type)
		}
		if data.Component != "" {
			db = db.Where(data.HasTableName("component")+" LIKE ?", "%"+data.Component+"%")
		}
		if data.Path != "" {
			db = db.Where(data.HasTableName("path")+" LIKE ?", "%"+data.Path+"%")
		}
		if data.Auth != "" {
			db = db.Where(data.HasTableName("api")+" LIKE ?", "%"+data.Auth+"%")
		}
		if data.Locale != "" {
			db = db.Where(data.HasTableName("locale")+" LIKE ?", "%"+data.Locale+"%")
		}
		if data.Icon != "" {
			db = db.Where(data.HasTableName("icon")+" LIKE ?", "%"+data.Icon+"%")
		}

		if data.Status != 0 {
			db = db.Where(data.HasTableName("status")+" = ?", data.Status)
		}
		return db
	}
}

func (data AuthRule) GetId() int64 {
	return data.Id
}

func (data AuthRule) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return validation.ValidateStruct(&data,
		validation.Field(&data.ParentId, validation.Min(0).Error("上级菜单不合法")),
		validation.Field(&data.Type,
			validation.Required.Error("菜单类型必须选择"),
			validation.In(RuleTypesMenu, RuleTypesButton, RuleTypesPage).Error("菜单类型不合法"),
		),
		validation.Field(&data.Locale, validation.Required.Error("菜单名称必填")),
		// validation.Field(&data.Path, validation.Required.Error("路由地址必填")),
		validation.Field(&data.Order, validation.Required.Error("显示排序必填")),
		validation.Field(&data.Status, validation.Required.Error("菜单状态必选")),
	)
}

func (data AuthRule) HasChildren() (childNums int64, err error) {
	err = db.DB().Model(&AuthRole{}).Where("id = ?", data.Id).Count(&childNums).Error
	return
}

func (data AuthRule) GetParentsNameWithRoute(route string) string {
	if route == "" {
		return ""
	}

	var sonAuth AuthRule
	if err := db.DB().Model(&sonAuth).Where("api_route = ?", route).First(&sonAuth).Error; err != nil {
		return ""
	}
	return data.GetParentsNameWithAuth(sonAuth.Auth)
}

func (data AuthRule) GetParentsNameWithAuth(auth string) string {
	if auth != "" {
		var auths []string
		for index, val := range strings.Split(auth, ":") {
			if index == 0 {
				auths = append(auths, val)
			} else {
				//auths = append(auths, utils.Strings(auths[index-1], ":", val).String())
				auths = append(auths, six.Str(auths[index-1]).Append(":", val).String())
			}
		}
		var list []string
		if err := db.DB().Table(data.TableName()).Where("auth in ?", auths).Select("locale").Order("id asc").Scan(&list).Error; err != nil {
			log.Println(err)
			return ""
		}
		return strings.Join(list, "--")
	}
	return ""
}
