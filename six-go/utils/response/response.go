package response

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var EmptyList = make([]map[string]any, 0)
var EmptyMap = make(map[string]any)

const (
	CodeOk     = 0
	CodeWarn   = 1
	CodeError  = 2
	CodeNoAuth = 3
	CodeLogout = 9
)

const (
	ErrorParamsParser       = "参数解析错误"
	ErrorDbNotFound         = "数据不存在"
	ErrorDbSelect           = "查询失败"
	ErrorDbCreate           = "添加失败"
	ErrorDbUpdate           = "修改失败"
	ErrorDbDelete           = "删除失败"
	ErrorSetting            = "设置失败"
	ErrorSystem             = "系统错误"
	ErrorSystemMiddle       = "系统错误：middleware"
	ErrorUsernameOrPassword = "用户名或密码错误"
	ErrorLogin              = "登录失败"
)

const (
	OkDbSelect = "查询成功"
	OkDbCreate = "添加成功"
	OkDbUpdate = "修改成功"
	OkDbDelete = "删除成功"
	OkSetting  = "设置成功"
	OkLogin    = "登录成功"
)

func Json(c *gin.Context, msg string, data any) {
	var result Result
	result.Code = CodeOk
	result.Msg = msg
	result.Data = data
	c.JSON(200, result)
}

func JsonPage(c *gin.Context, msg string, data any, total int64) {
	var result Result
	result.Code = CodeOk
	result.Msg = msg
	result.Data = map[string]any{
		"total": total,
		"list":  data,
	}
	c.JSON(200, result)
}

func JsonWarn(c *gin.Context, msg string) {
	var result Result
	result.Code = CodeWarn
	result.Msg = msg
	c.JSON(200, result)
}

func JsonError(c *gin.Context, msg string, errno error) {
	log.Println(errno)
	var result Result
	result.Code = CodeError
	result.Msg = msg
	c.JSON(200, result)
}

func JsonLogout(c *gin.Context) {
	var result Result
	result.Code = CodeLogout
	result.Msg = "登录状态失效"
	c.AbortWithStatusJSON(http.StatusOK, result)
}

func JsonNoAuth(c *gin.Context) {
	var result Result
	result.Code = CodeNoAuth
	result.Msg = "没有操作权限"
	c.AbortWithStatusJSON(http.StatusOK, result)
}
