package dao

import (
	"github.com/vainback/six-util/v3"
	"gorm.io/gorm"
)

type QueryKeyword struct {
	Keyword string `json:"keyword"`
}

func (ts QueryKeyword) SqlBuilder(fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ts.Keyword) == 0 {
			return db
		}

		sql := six.Strings(fields[0], " like '%", ts.Keyword, "%'")
		if len(fields) > 1 {
			for _, field := range fields[1:] {
				sql.Append(" or ", field, " like '%", ts.Keyword, "%'")
			}
		}
		return db.Where(sql.String())
	}
}
