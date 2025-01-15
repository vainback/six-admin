package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) Dict(app *gin.RouterGroup) {
	p := app.Group("/dict")
	p.POST("/list", controller.Dict.List)
	p.POST("/select", controller.Dict.Select)
	p.POST("/get", controller.Dict.Get)
	p.POST("/add", controller.Dict.Add)
	p.POST("/edit", controller.Dict.Update)
	p.POST("/save", controller.Dict.Save)
	p.POST("/del", controller.Dict.Delete)
}
