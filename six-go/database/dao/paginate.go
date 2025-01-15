package dao

import "gorm.io/gorm"

type QueryPage struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func (ts QueryPage) SqlBuilder() func(db *gorm.DB) *gorm.DB {
	if ts.Page <= 0 {
		ts.Page = 1
	}

	if ts.Limit <= 0 {
		ts.Limit = 20
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((ts.Page - 1) * ts.Limit).Limit(ts.Limit)
	}
}
