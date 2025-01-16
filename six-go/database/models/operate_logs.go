package models

import (
	"github.com/vainback/six-util/v3"
	"strings"

	"gorm.io/gorm"
	"six-go/database/db"
)

var TableOperateLog OperateLog

type OperateLog struct {
	db.MODEL
	db.SoftDelete
	TenantIdentify
	IP           string `json:"ip"`                                 // 客户端IP
	Uid          int64  `json:"uid"`                                // 请求的用户
	Route        string `json:"route"`                              // 请求的路由
	RouteName    string `json:"route_name"`                         // 路由名
	RequestBody  string `json:"request_body" gorm:"type:longtext"`  // 请求body
	ResponseBody string `json:"response_body" gorm:"type:longtext"` // 响应body
	Latency      string `json:"latency"`
	Agent        string `json:"agent"`
	Method       string `json:"method"`
}

func (OperateLog) TableName() string {
	return "operate_logs"
}

func (data OperateLog) HasTableName(filed string) string {
	return six.Strings(data.TableName(), ".", filed).String()
}

func (data OperateLog) KeywordFields() []string {
	return six.Arrays(
		data.HasTableName("ip"),
		data.HasTableName("route"),
		data.HasTableName("route_name"),
		data.HasTableName("request_body"),
		data.HasTableName("response_body"),
	)
}

func (data OperateLog) FilterSqlBuilder() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if data.Id > 0 {
			return db.Where(data.HasTableName("id")+" = ?", data.Id)
		}
		if data.IP != "" {
			db = db.Where(data.HasTableName("ip")+" LIKE ?", "%"+data.IP+"%")
		}
		if data.Uid != 0 {
			db = db.Where(data.HasTableName("uid")+" = ?", data.Uid)
		}
		if data.Route != "" {
			db = db.Where(data.HasTableName("route")+" LIKE ?", "%"+data.Route+"%")
		}
		if data.RouteName != "" {
			db = db.Where(data.HasTableName("route_name")+" LIKE ?", "%"+data.RouteName+"%")
		}
		if data.RequestBody != "" {
			db = db.Where(data.HasTableName("request_body")+" LIKE ?", "%"+data.RequestBody+"%")
		}
		if data.ResponseBody != "" {
			db = db.Where(data.HasTableName("response_body")+" LIKE ?", "%"+data.ResponseBody+"%")
		}
		return db.Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))
	}
}

func (data OperateLog) GetId() int64 {
	return data.Id
}
