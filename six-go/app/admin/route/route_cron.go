package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) Cron(app *gin.RouterGroup) {
	p := app.Group("/cron")
	p.POST("/list", controller.Cron.List)
	p.POST("/select", controller.Cron.Select)
	p.POST("/get", controller.Cron.Get)
	p.POST("/add", controller.Cron.Add)
	p.POST("/edit", controller.Cron.Update)
	p.POST("/save", controller.Cron.Save)
	p.POST("/del", controller.Cron.Delete)
}
