package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) Files(app *gin.RouterGroup) {
	p := app.Group("/files")
	p.POST("/list", controller.Files.List)
	p.POST("/select", controller.Files.Select)
	p.POST("/get", controller.Files.Get)
	p.POST("/edit", controller.Files.Update)
	p.POST("/save", controller.Files.Save)
	p.POST("/del", controller.Files.Delete)
}
