package dao

import (
	"six-go/utils"
	"strings"

	"gorm.io/gorm"
)

type QueryTime struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (ts QueryTime) SqlBuilder(tableName ...string) func(db *gorm.DB) *gorm.DB {
	var field = "create_time"
	if len(tableName) > 0 {
		//field = fields[0]
		field = utils.String(tableName[0]).Strings(".", field).String()
	}

	return func(db *gorm.DB) *gorm.DB {
		if start := strings.TrimSpace(ts.StartTime); start != "" {
			db = db.Where(field+" > ?", start)
		}

		if end := strings.TrimSpace(ts.EndTime); end != "" {
			db = db.Where(field+" < ?", end)
		}

		return db
	}
}
