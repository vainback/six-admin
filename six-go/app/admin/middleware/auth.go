package middleware

import (
	authService "six-go/app/admin/service/auth"
	"strings"

	"github.com/gin-gonic/gin"
	"six-go/utils/response"
)

func Authorization(c *gin.Context) {
	user := SessionUser(c)
	if !user.IsRoot {
		if len(user.RoleIds) == 0 {
			c.Abort()
			response.JsonNoAuth(c)
			return
		}

		route := strings.ReplaceAll(c.Request.URL.Path, "/admin/", "/")
		if has := authService.HasAuth(route, user.RoleIds...); !has {
			c.Abort()
			response.JsonNoAuth(c)
			return
		}
		/*var ok int64
		if err := db.DB().Table(models.TableAuthRule.TableName()+" as rule").
			Joins(fmt.Sprintf("left join %s as relation on relation.rule_id = rule.id", models.TableAuthRelation.TableName())).
			Joins(fmt.Sprintf("left join %s as role on role.id = relation.role_id", models.TableAuthRole.TableName())).
			Select("rule.status as rule_status,role.status as role_status").
			Where("rule.api_route = ?", route).
			Where("role.id in ?", user.RoleIds).
			Where("rule.status = ?", true).
			Where("role.status = ?", true).
			Count(&ok).Error; err != nil {
			c.Abort()
			log.Println(err)
			response.JsonNoAuth(c)
			return
		}
		if ok == 0 {
			c.Abort()
			response.JsonNoAuth(c)
			return
		}*/
	}
	c.Next()
}
