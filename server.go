package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuritaeiji/gin-practice/controller"
	"github.com/kuritaeiji/gin-practice/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		videoController.Index(ctx)
	})
	server.POST("/videos", func(ctx *gin.Context) {
		videoController.Create(ctx)
	})

	server.Run(":8080")
}
