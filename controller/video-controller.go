package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kuritaeiji/gin-practice/model"
	"github.com/kuritaeiji/gin-practice/service"
)

type VideoController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service,
	}
}

func (c *controller) Index(ctx *gin.Context) {
	videos := c.service.FindAll()
	ctx.JSON(200, videos)
}

func (c *controller) Create(ctx *gin.Context) {
	var video model.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.service.Save(video)
	ctx.JSON(200, video)
}
