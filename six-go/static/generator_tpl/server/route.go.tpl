package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) <Tpl-Module>(app *gin.RouterGroup) {
    p := app.Group("<Tpl-Route-Group>")
    p.POST("/list", controller.<Tpl-Module>.List)
    p.POST("/tree-select", controller.<Tpl-Module>.TreeSelect)
    p.POST("/select", controller.<Tpl-Module>.Select)
    p.POST("/get", controller.<Tpl-Module>.Get)
    p.POST("/add", controller.<Tpl-Module>.Add)
    p.POST("/edit", controller.<Tpl-Module>.Update)
    p.POST("/save", controller.<Tpl-Module>.Save)
    p.POST("/del", controller.<Tpl-Module>.Delete)
}