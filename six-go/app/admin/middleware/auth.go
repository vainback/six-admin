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
	}
	c.Next()
}
