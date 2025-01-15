package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) Organization(app *gin.RouterGroup) {
	p := app.Group("/organization")
	p.POST("/list", controller.Organization.List)
	p.POST("/tree-select", controller.Organization.TreeSelect)
	p.POST("/get", controller.Organization.Get)
	p.POST("/add", controller.Organization.Add)
	p.POST("/edit", controller.Organization.Update)
	p.POST("/save", controller.Organization.Save)
	p.POST("/del", controller.Organization.Delete)
}
