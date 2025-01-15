package adminRoute

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"six-go/app/admin/middleware"
)

// LogMode 不记录Log
//const LogMode = middleware.LogModeClosed

// LogMode 全量记录 几乎完全记录请求和响应body
const LogMode = middleware.LogModeFull

// LogMode 简单记录 不记录请求和响应body
//const LogMode = middleware.LogModeSimple

type Route struct {
	auths map[string]string
}

type SingleRoute struct{}

var singleRoute SingleRoute

func Routers(app *gin.RouterGroup) {
	var route = Route{}

	singleRoute.AuthUserSingle(app)

	app.Use(middleware.IsLogin)
	app.Use(middleware.TenantDI)

	app.Use(middleware.Logs(LogMode))

	app.Use(middleware.Authorization)

	routeType := reflect.TypeOf(route)
	routeTypeValue := reflect.ValueOf(route)
	appValue := reflect.ValueOf(app)
	for i := 0; i < routeType.NumMethod(); i++ {
		methodValue := routeTypeValue.Method(i)
		methodValue.Call([]reflect.Value{appValue})
	}
}
