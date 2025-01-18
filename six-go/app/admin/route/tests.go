package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) Tests(app *gin.RouterGroup) {
	p := app.Group("tests")
	p.POST("/list", controller.Tests.List)
	p.POST("/select", controller.Tests.Select)
	p.POST("/get", controller.Tests.Get)
	p.POST("/add", controller.Tests.Add)
	p.POST("/edit", controller.Tests.Update)
	p.POST("/save", controller.Tests.Save)
	p.POST("/del", controller.Tests.Delete)
}
