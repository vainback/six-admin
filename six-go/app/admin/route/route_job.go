package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) Job(app *gin.RouterGroup) {
	p := app.Group("/job")
	p.POST("/list", controller.Job.List)
	p.POST("/select", controller.Job.Select)
	p.POST("/tree-select", controller.Job.TreeSelect)
	p.POST("/get", controller.Job.Get)
	p.POST("/add", controller.Job.Add)
	p.POST("/edit", controller.Job.Update)
	p.POST("/save", controller.Job.Save)
	p.POST("/del", controller.Job.Delete)
}
