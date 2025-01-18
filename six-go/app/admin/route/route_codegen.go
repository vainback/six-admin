package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) Codegen(app *gin.RouterGroup) {
	p := app.Group("/codegen")
	p.POST("/list", controller.CodeGenerator.List)
	p.POST("/add", controller.CodeGenerator.Add)
	p.POST("/save", controller.CodeGenerator.Save)
	p.POST("/del", controller.CodeGenerator.Delete)
	p.POST("/generator", controller.CodeGenerator.Generator)
}
