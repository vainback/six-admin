package main

import (
	"log"
	"net/http"
	"six-go/extra/crons"

	"github.com/gin-gonic/gin"
	adminRoute "six-go/app/admin/route"
	"six-go/config"
	"six-go/database/dao"
	"six-go/database/db"
)

func init() {
	go func() {
		if config.DB().AutoMigrate() {
			dao.AutoMigrate()
		}

		// 自动拉取数据库定时任务 自动 添加到定时服务中
		crons.Init()

		select {
		case <-config.Notices:
			db.Reconnect()
		}
	}()

}

func main() {
	if config.APP().Debug() {
		gin.SetMode(gin.DebugMode)
	}
	app := gin.Default()
	app.Use(CORS)

	app.StaticFS("uploads", http.Dir("./uploads"))

	app.LoadHTMLFiles("dist/index.html")
	app.StaticFS("assets", http.Dir("./dist/assets"))
	app.GET("/admin/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	adminRoute.Routers(app.Group("/admin"))

	log.Fatalln(app.Run(config.APP().ListenPort()))
}

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
