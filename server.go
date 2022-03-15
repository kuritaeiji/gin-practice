package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuritaeiji/gin-practice/controller"
	"github.com/kuritaeiji/gin-practice/middleware"
	"github.com/kuritaeiji/gin-practice/service"
	"github.com/kuritaeiji/gin-practice/utils"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	utils.SettingLogFile()
	utils.RegisterValidators()

	server.Use(middleware.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		videoController.Index(ctx)
	})
	server.POST("/videos", func(ctx *gin.Context) {
		videoController.Create(ctx)
	})

	server.Run(":8080")
}
