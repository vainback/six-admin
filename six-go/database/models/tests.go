package models

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vainback/six-util/v3"
	"gorm.io/gorm"
	"six-go/database/db"
)

var TableTests Tests

type Tests struct {
	db.MODEL
	db.SoftDelete
	TenantIdentify
	TeStr            string   `json:"te_str" gorm:"type:varchar(255);column:te_str;comment:测试普通字符串;not null;default:'';index;"`
	TeInt            int64    `json:"te_int" gorm:"type:int(11);column:te_int;comment:测试普通整形;not null;default:0;index;"`
	TeBigint         int64    `json:"te_bigint" gorm:"type:bigint(20);column:te_bigint;comment:测试大整形;not null;default:0;index;"`
	TeFloat          float64  `json:"te_float" gorm:"type:float;column:te_float;comment:测试浮点数;not null;default:0.00;"`
	TeDecimal        float64  `json:"te_decimal" gorm:"type:decimal(10,2);column:te_decimal;comment:测试高精度浮点;not null;default:0.00;"`
	TeSelect         string   `json:"te_select" gorm:"type:varchar(32);column:te_select;comment:测试单选下拉框;not null;default:'';"`
	TeSelectMany     []string `json:"te_select_many" gorm:"type:text;serializer:json;column:te_select_many;comment:测试多选下拉框;"`
	TeRadio          string   `json:"te_radio" gorm:"type:varchar(32);column:te_radio;comment:单选框;not null;default:'';"`
	TeCheckbox       []string `json:"te_checkbox" gorm:"type:text;serializer:json;column:te_checkbox;comment:复选框;"`
	TeSwitch         bool     `json:"te_switch" gorm:"type:tinyint(1);column:te_switch;comment:开关;not null;default:0;"`
	TeTimepicker     string   `json:"te_timepicker" gorm:"type:varchar(32);column:te_timepicker;comment:时间选择器;"`
	TeDatepicker     string   `json:"te_datepicker" gorm:"type:varchar(32);column:te_datepicker;comment:日期选择器;"`
	TeDatetimepicker string   `json:"te_datetimepicker" gorm:"type:varchar(32);column:te_datetimepicker;comment:日期时间选择器;"`
	TeImageOne       []string `json:"te_image_one" gorm:"type:text;serializer:json;column:te_image_one;comment:上传图片;{{tpl - $index}}"`
	TeImageMany      []string `json:"te_image_many" gorm:"type:text;serializer:json;column:te_image_many;comment:上传图片（多图）;{{tpl - $index}}"`
	TeVideo          []string `json:"te_video" gorm:"type:text;serializer:json;column:te_video;comment:上传视频;{{tpl - $index}}"`
	TeFile           []string `json:"te_file" gorm:"type:text;serializer:json;column:te_file;comment:上传文件;{{tpl - $index}}"`
	TeText           string   `json:"te_text" gorm:"type:text;column:te_text;comment:文本域;"`
	TeEditor         string   `json:"te_editor" gorm:"type:longtext;column:te_editor;comment:富文本;"`
}

func (data Tests) TableName() string {
	return "tests"
}

func (data Tests) HasTableName(filed string) string {
	return six.Strings(data.TableName(), ".", filed).String()
}

func (data Tests) KeywordFields() []string {
	return six.Arrays(
		data.HasTableName("te_str"),
		data.HasTableName("te_int"),
		data.HasTableName("te_bigint"),
	)
}

func (data Tests) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))
	}
}

func (data Tests) GetId() int64 {
	return data.Id
}

func (data Tests) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return validation.ValidateStruct(&data,
		validation.Field(&data.TeStr, validation.Required.Error("测试普通字符串不能为空")),
		validation.Field(&data.TeInt, validation.Required.Error("测试普通整形不能为空")),
		validation.Field(&data.TeBigint, validation.Required.Error("测试大整形不能为空")),
		validation.Field(&data.TeFloat, validation.Required.Error("测试浮点数不能为空")),
		validation.Field(&data.TeDecimal, validation.Required.Error("测试高精度浮点不能为空")),
		validation.Field(&data.TeSelect, validation.Required.Error("测试单选下拉框不能为空")),
		validation.Field(&data.TeSelectMany, validation.Required.Error("测试多选下拉框不能为空")),
		validation.Field(&data.TeRadio, validation.Required.Error("单选框不能为空")),
		validation.Field(&data.TeCheckbox, validation.Required.Error("复选框不能为空")),
		validation.Field(&data.TeSwitch, validation.Required.Error("开关不能为空")),
		validation.Field(&data.TeTimepicker, validation.Required.Error("时间选择器不能为空")),
		validation.Field(&data.TeDatepicker, validation.Required.Error("日期选择器不能为空")),
		validation.Field(&data.TeDatetimepicker, validation.Required.Error("日期时间选择器不能为空")),
		validation.Field(&data.TeImageOne, validation.Required.Error("上传图片不能为空")),
		validation.Field(&data.TeImageMany, validation.Required.Error("上传图片（多图）不能为空")),
		validation.Field(&data.TeFile, validation.Required.Error("上传文件不能为空")),
		validation.Field(&data.TeText, validation.Required.Error("文本域不能为空")),
		validation.Field(&data.TeEditor, validation.Required.Error("富文本不能为空")),
	)
}
