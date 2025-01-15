package dao

import (
	"strings"

	"gorm.io/gorm"
)

type QueryOrderBys struct {
	OrderBy []QueryOrderBy `json:"order_by"`
}

type QueryOrderBy struct {
	Field  string `json:"field"`
	IsDesc bool   `json:"is_desc"`
}

const OrderByDefault = "create_time desc"

func (ts QueryOrderBys) SqlBuilder() func(db *gorm.DB) *gorm.DB {
	var isDesc = map[bool]string{
		false: "asc",
		true:  "desc",
	}
	return func(db *gorm.DB) *gorm.DB {
		obl := len(ts.OrderBy)
		if obl == 0 {
			return db.Order(OrderByDefault)
		}

		var sql strings.Builder
		for _, orderBy := range ts.OrderBy {
			sql.WriteString(ts.OrderBy[0].Field)
			sql.WriteString(" ")
			sql.WriteString(isDesc[orderBy.IsDesc])
			sql.WriteString(",")
		}
		orderBySql := strings.Trim(sql.String(), ",")
		return db.Order(orderBySql)
	}
}
