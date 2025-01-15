package adminRoute

import (
	"github.com/gin-gonic/gin"
	"six-go/app/admin/controller"
)

func (r Route) AuthRule(app *gin.RouterGroup) {
	p := app.Group("/auth/rule")
	p.POST("/list", controller.AuthRule.List)
	p.POST("/tree-select", controller.AuthRule.TreeSelect)
	p.POST("/select", controller.AuthRule.Select)
	p.POST("/get", controller.AuthRule.Get)
	p.POST("/add", controller.AuthRule.Add)
	p.POST("/edit", controller.AuthRule.Update)
	p.POST("/save", controller.AuthRule.Save)
	p.POST("/del", controller.AuthRule.Delete)
}

func (r Route) AuthRole(app *gin.RouterGroup) {
	p := app.Group("/auth/role")
	p.POST("/list", controller.AuthRole.List)
	p.POST("/select", controller.AuthRole.Select)
	p.POST("/tree-select", controller.AuthRole.TreeSelect)
	p.POST("/get", controller.AuthRole.Get)
	p.POST("/add", controller.AuthRole.Add)
	p.POST("/edit", controller.AuthRole.Update)
	p.POST("/save", controller.AuthRole.Save)
	p.POST("/del", controller.AuthRole.Delete)
}

func (r Route) AuthRelation(app *gin.RouterGroup) {
	p := app.Group("/auth/relation")
	p.POST("/select/group/role", controller.AuthRelation.SelectGroupRole) // map[role_id][]{rule_ids} 角色、拥有哪些规则
	p.POST("/select/group/rule", controller.AuthRelation.SelectGroupRule) // map[rule_id][]{role_ids} 规则、属于哪些角色
	p.POST("/select/rule", controller.AuthRelation.SelectRuleIdsWithRoleId)
	p.POST("/set", controller.AuthRelation.Set)
}

func (r Route) AuthUser(app *gin.RouterGroup) {
	p := app.Group("/auth/user")
	p.POST("/list", controller.AuthUser.List)
	p.POST("/select", controller.AuthUser.Select)
	p.POST("/get", controller.AuthUser.Get)
	p.POST("/add", controller.AuthUser.Add)
	p.POST("/edit", controller.AuthUser.Update)
	p.POST("/save", controller.AuthUser.Save)
	p.POST("/del", controller.AuthUser.Delete)
}
