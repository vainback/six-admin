package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) Tenant(app *gin.RouterGroup) {
	p := app.Group("/tenant")
	p.POST("/list", controller.Tenant.List)
	p.POST("/select", controller.Tenant.Select)
	p.POST("/get", controller.Tenant.Get)
	p.POST("/add", controller.Tenant.Add)
	p.POST("/edit", controller.Tenant.Update)
	p.POST("/save", controller.Tenant.Save)
	p.POST("/del", controller.Tenant.Delete)
}
