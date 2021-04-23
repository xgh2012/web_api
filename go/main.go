package main

import (
	"demo/api"
	"demo/api/controlls"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	Version = "1.0.1"
	Time    = "2021年3月2日14:35:31 "
	GitLog  = "719c50c16d045ce207cfbb02a2a3047eb09591bd"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	RegistRouter(router)

	//router.StaticFS("/web", http.Dir("/Users/xhome/web_api/html"))
	router.StaticFS("/web", http.Dir("F:/Git_Code/docker/web_api/html"))

	log.Println(Version, Time, GitLog)
	log.Println("启动后端服务，端口：")

	err1 := router.Run(":8050")
	log.Fatalln("启动后端服务失败：", err1)
}

func RegistRouter(router *gin.Engine) {
	router.GET("/api/ok", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	new(controlls.System).Router(router)
	new(api.Login).Router(router)
	new(api.Index).Router(router)
}
