package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
	"six-go/app/admin/middleware"
)

func (r SingleRoute) AuthUserSingle(app *gin.RouterGroup) {
	p := app.Group("/user/single")
	p.POST("/tenant", controller.UserSingle.GetTenant)
	p.POST("/login", controller.UserSingle.Login)

	p.Use(middleware.IsLogin)
	p.POST("/info", controller.UserSingle.Userinfo)
	p.POST("/logout", controller.UserSingle.Logout)
	p.POST("/menu", middleware.TenantDI, controller.UserSingle.Menu)
	p.POST("/btns", controller.UserSingle.Btns)

	p.Use(middleware.Logs(LogMode))
	p.POST("/update", controller.UserSingle.Update)
	p.POST("/reset/password", controller.UserSingle.ResetPassword)
	p.POST("/upload", controller.UserSingle.Upload)
}
