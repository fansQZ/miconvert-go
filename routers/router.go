// Package routers
// @Author: fzw
// @Create: 2022/10/8
// @Description: 路由相关
package routers

import (
	"github.com/gin-gonic/gin"
	"miconvert-go/controllers"
	"miconvert-go/interceptor"
	"miconvert-go/setting"
	"miconvert-go/ws"
)

//
// Run
//  @Description: 启动路由
//
func Run() {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	//设置静态文件位置
	r.Static("/static", "/")
	//ping
	r.GET("/ping", controllers.Ping)
	//游客解析
	convert := r.Group("/convert")
	{
		convertController := controllers.NewConvertController()
		convert.POST("/convertFile", convertController.ConvertFile)
		convert.GET("/getSupportFormat", convertController.GetSupportOutFormat)
		convert.GET("/downloadFile/:filename", convertController.DownloadFile)
	}
	//添加token拦截器
	r.Use(interceptor.TokenAuthorize())
	//ws
	r.GET("/ws", func(ctx *gin.Context) {
		ws.ServeWs(ctx.Writer, ctx.Request)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
