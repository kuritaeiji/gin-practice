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
	loginController controller.AuthController  = controller.NewAuthController()
)

func main() {
	server := gin.Default()

	utils.SettingLogFile()
	utils.RegisterValidators()

	notAuth := server.Group("")
	{
		notAuth.POST("/login", loginController.Login)
	}

	server.Use(middleware.Logger(), middleware.AuthorizeJWT())

	auth := server.Group("")
	{
		auth.GET("/videos", func(ctx *gin.Context) {
			videoController.Index(ctx)
		})
		auth.POST("/videos", func(ctx *gin.Context) {
			videoController.Create(ctx)
		})
	}
	auth.Use(middleware.AuthorizeJWT())

	server.Run(":8080")
}
