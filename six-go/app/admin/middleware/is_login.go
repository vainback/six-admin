package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	loginService "six-go/app/admin/service/login"
	"six-go/database/models"
	"six-go/utils/response"
)

func IsLogin(c *gin.Context) {
	token := GetHeaderToken(c)
	if len(token) == 0 {
		response.JsonLogout(c)
		return
	}

	sessionStorage := loginService.New(token, true).Get()
	if !sessionStorage.IsLogin {
		response.JsonLogout(c)
		return
	}
	c.Set(SessionUserKey, sessionStorage.Userinfo)
	c.Next()
}

const SessionUserKey = "userinfo"
const HeaderTokenKey = "Six-Token"

func SessionUser(c *gin.Context) *models.AuthUser {
	value, ok := c.Get(SessionUserKey)
	if !ok || value == nil {
		return nil
	}
	return value.(*models.AuthUser)
}

func GetHeaderToken(c *gin.Context) string {
	return strings.TrimSpace(c.GetHeader(HeaderTokenKey))
}
