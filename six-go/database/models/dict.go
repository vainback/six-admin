package models

import (
	"errors"
	"github.com/vainback/six-util/v3"
	"six-go/database/v_errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"six-go/database/db"
)

var TableDict Dict

type Dict struct {
	db.MODEL
	db.SoftDelete
	TenantIdentify
	Type   string `json:"type" gorm:"index"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Color  string `json:"color"`
	IsSync int    `json:"is_sync" gorm:"type:tinyint(1);default:-1;comment:状态 -1 不同步 1同步"` // 是否同步 开启同步的数据 在添加时将自动同步给全部租户 此功能仅限root租户操作 当CURD条件不启用租户ID时，该字段不生效且只有root租户可以对字典增删改
}

func (data Dict) TableName() string {
	return "dict"
}

func (data Dict) HasTableName(filed string) string {
	return six.Strings(data.TableName(), ".", filed).String()
}

func (data Dict) KeywordFields() []string {
	return six.Arrays(
		data.HasTableName("type"),
		data.HasTableName("label"),
	)
}

func (data Dict) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.Type != "" {
			db = db.Where(data.HasTableName("type")+" = ?", data.Type)
		}
		if data.Label != "" {
			db = db.Where(data.HasTableName("label")+" LIKE ?", "%"+data.Label+"%")
		}

		return db.Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))
	}
}

func (data Dict) GetId() int64 {
	return data.Id
}

func (data Dict) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return validation.ValidateStruct(&data,
		validation.Field(&data.Type, validation.Required.Error("字典类型不能为空")),
		validation.Field(&data.Label, validation.Required.Error("字典名不能为空")),
		validation.Field(&data.Value, validation.Required.Error("字典值不能为空")),
	)
}

func (data Dict) CheckUnique() error {
	var row Dict
	if err := db.DB().Where("type = ?", data.Type).Where("value = ?", data.Value).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	if row.Id != data.Id {
		return v_errors.DictUnique
	}
	return nil
}

func (data Dict) HasChildren() (childNums int64, err error) {
	err = db.DB().Model(&Dict{}).Where("type = ?", data.Value).Count(&childNums).Error
	return
}
