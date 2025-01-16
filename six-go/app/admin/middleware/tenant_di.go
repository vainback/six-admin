package middleware

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"six-go/utils/response"
)

// TenantDI 对request body params 校验/注入 session user tenant_id 参数
func TenantDI(c *gin.Context) {
	requestBodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Abort()
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}
	var requestBody map[string]any
	if err = jsoniter.Unmarshal(requestBodyBytes, &requestBody); err != nil {
		c.Abort()
		response.JsonError(c, response.ErrorParamsParser, err)
		return
	}

	sessionUser := SessionUser(c)

	requestBodyModel, ok := requestBody["model"]
	if ok {
		m := cast.ToStringMap(requestBodyModel)
		if tenantId, tok := m["tenant_id"]; tok {
			if cast.ToInt64(m["id"]) > 0 && cast.ToInt64(tenantId) != sessionUser.TenantId {
				c.Abort()
				response.JsonError(c, "操作被禁止", nil)
				return
			}
		}
		m["tenant_id"] = sessionUser.TenantId
		requestBody["model"] = m
	}

	bodyBytes, err := jsoniter.Marshal(requestBody)
	if err != nil {
		c.Abort()
		response.JsonError(c, "参数序列化失败", err)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	c.Next()
}
