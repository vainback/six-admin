package entity

import (
	"fmt"
	"six-go/utils"
	"strings"

	"gorm.io/gorm"
	"six-go/database/models"
)

type AuthUserJoinJob struct {
	models.AuthUser
	Job models.Job `json:"job" gorm:"embedded;embeddedPrefix:job_"`
}

func (e AuthUserJoinJob) TableName() string {
	return e.AuthUser.TableName()
}

func (e AuthUserJoinJob) KeywordFields() []string {
	return append(e.AuthUser.KeywordFields(), e.Job.KeywordFields()...)
}

func (e AuthUserJoinJob) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(e.AuthUser.FilterSqlBuilder()).Scopes(e.Job.FilterSqlBuilder())
	}
}

func (e AuthUserJoinJob) Fields() string {
	var fields = []string{
		utils.Strings(e.AuthUser.TableName(), ".*").String(),
		utils.Strings(e.Job.TableName(), ".name as job_name").String(),
	}
	return strings.Join(fields, ",")
}

func (e AuthUserJoinJob) LeftJoin() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(e.AuthUser.TableName()).
			Joins(fmt.Sprintf("LEFT JOIN %s ON %s.id = %s.job_id", e.Job.TableName(), e.Job.TableName(), e.AuthUser.TableName()))
	}
}

func (e AuthUserJoinJob) Select() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(e.Fields())
	}
}
