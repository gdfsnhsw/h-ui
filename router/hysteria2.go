package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initHysteria2AuthRouter(hysteria2Api *gin.RouterGroup) {
	hysteria2 := hysteria2Api.Group("/hysteria2")
	{
		hysteria2.POST("/auth", controller.Hysteria2Auth)
	}
}

func initHysteria2Router(hysteria2Api *gin.RouterGroup) {
	hysteria2 := hysteria2Api.Group("/hysteria2")
	{
		hysteria2.POST("/startHysteria2", controller.StartHysteria2)
		hysteria2.POST("/stopHysteria2", controller.StopHysteria2)
		hysteria2.POST("/restartHysteria2", controller.RestartHysteria2)
		hysteria2.POST("/hysteria2Kick", controller.Hysteria2Kick)
		hysteria2.GET("/getHysteria2Version", controller.GetHysteria2Version)
	}
}
