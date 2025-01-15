package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) OperateLogs(app *gin.RouterGroup) {
	p := app.Group("/logs")
	p.POST("/list", controller.OperateLogs.List)
	p.POST("/select", controller.OperateLogs.Select)
	p.POST("/get", controller.OperateLogs.Get)
}
