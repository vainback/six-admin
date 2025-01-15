package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	authService "six-go/app/admin/service/auth"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"six-go/database/db"
	"six-go/database/models"
)

type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.b.Write(b)
	return w.ResponseWriter.Write(b)
}

var logMode = map[mode]func(c *gin.Context){
	LogModeSimple: SimpleLog,
	LogModeFull:   OperateLog,
	LogModeClosed: ClosedLog,
}

type mode int

const (
	LogModeClosed mode = 0
	LogModeSimple mode = 1
	LogModeFull   mode = 9
)

func Logs(m mode) func(c *gin.Context) {
	return logMode[m]
}

func ClosedLog(c *gin.Context) {
	c.Next()
}

func SimpleLog(c *gin.Context) {
	now := time.Now()
	route := strings.ReplaceAll(c.Request.URL.Path, "/admin/", "/")
	var myLog = models.OperateLog{
		IP:        c.ClientIP(),
		Route:     route,
		RouteName: authService.GetNamesByRoute(route),
		Agent:     c.Request.UserAgent(),
		Method:    c.Request.Method,
	}
	sessionUser := SessionUser(c)
	if sessionUser != nil {
		myLog.Uid = sessionUser.Id
		myLog.TenantId = sessionUser.TenantId
	}
	c.Next()
	myLog.Latency = fmt.Sprintf("%.3f秒", time.Since(now).Seconds())
	go func() {
		if err := db.DB().Create(&myLog).Error; err != nil {
			log.Printf("---- 日志记录出错 ----\n  时间：%s\n  错误信息：%s\n---- 日志记录出错 End ----\n", time.Now().String(), err.Error())
		}
	}()
}

func OperateLog(c *gin.Context) {
	now := time.Now()
	route := strings.ReplaceAll(c.Request.URL.Path, "/admin/", "/")
	var myLog = models.OperateLog{
		IP:        c.ClientIP(),
		Route:     route,
		RouteName: authService.GetNamesByRoute(route),
		Agent:     c.Request.UserAgent(),
		Method:    c.Request.Method,
	}

	sessionUser := SessionUser(c)
	if sessionUser != nil {
		myLog.Uid = sessionUser.Id
		myLog.TenantId = sessionUser.TenantId
	}

	if _, e := c.FormFile("file"); e != nil {
		requestBody, err := io.ReadAll(c.Request.Body)
		if err == nil {
			myLog.RequestBody = string(requestBody)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	}

	// 获取响应体
	responseW := responseWriter{
		c.Writer,
		bytes.NewBuffer([]byte{}),
	}
	c.Writer = responseW

	c.Next()
	myLog.ResponseBody = responseW.b.String()
	if len(myLog.ResponseBody) > 255 {
		//var result response.Result
		//_ = json.Unmarshal([]byte(myLog.ResponseBody), &result)
		//result.Data = nil
		//bys, _ := json.Marshal(result)
		myLog.ResponseBody = ""
	}
	myLog.Latency = fmt.Sprintf("%.3f秒", time.Since(now).Seconds())
	go func() {
		if err := db.DB().Create(&myLog).Error; err != nil {
			log.Printf("---- 日志记录出错 ----\n  时间：%s\n  错误信息：%s\n---- 日志记录出错 End ----\n", time.Now().String(), err.Error())
		}
	}()
}
