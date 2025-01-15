package dao

import "gorm.io/gorm"

type QuerySoftDelete struct {
	IsDelete bool `json:"is_delete"`
}

func (ts QuerySoftDelete) SqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !ts.IsDelete {
			return db
		}
		return db.Unscoped()
	}
}
