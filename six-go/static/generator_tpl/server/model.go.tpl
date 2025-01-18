package models

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vainback/six-util/v3"
	"gorm.io/gorm"
	"six-go/database/db"
)

var Table<Tpl-Module> <Tpl-Module>

type <Tpl-Module> struct {
	db.MODEL
	<Tpl-SoftDelete>
	<Tpl-Tenant>
	<Tpl-Fields>
}

func (data <Tpl-Module>) TableName() string {
	return "<Tpl-TableName>"
}

func (data <Tpl-Module>) HasTableName(filed string) string {
	return six.Strings(data.TableName(), ".", filed).String()
}

func (data <Tpl-Module>) KeywordFields() []string {
	return <Tpl-KeywordsFields>
}

func (data <Tpl-Module>) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db<Tpl-TenantWhere>
	}
}

func (data <Tpl-Module>) GetId() int64 {
	return data.Id
}

func (data <Tpl-Module>) Valid(idRequired ...bool) error {
	if len(idRequired) > 0 && idRequired[0] {
		if data.Id <= 0 {
			return errors.New("id is required")
		}
	}

	return <Tpl-ValidationRequired>
}

<Tpl - HasChildren>
func (data <Tpl-Module>) HasChildren() (childNums int64, err error) {
	err = db.DB().Model(&<Tpl-Module>{}).Where("parent_id = ?", data.Id).Count(&childNums).Error
	return
}
