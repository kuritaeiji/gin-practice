package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kuritaeiji/gin-practice/dto"
	"github.com/kuritaeiji/gin-practice/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewAuthController() *authController {
	return &authController{
		loginService: service.NewLoginService(),
		jwtService:   service.NewJWTService(),
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var credential dto.Credentials
	err := ctx.BindJSON(&credential)
	if err != nil {
		ctx.JSON(400, gin.H{"errors": err.Error()})
		return
	}
	if c.loginService.Login(credential.Username, credential.Password) {
		ts := c.jwtService.GenerateJWT(credential.Username, true)
		ctx.JSON(200, gin.H{
			"token": ts,
		})
		return
	}
	ctx.JSON(400, gin.H{"erros": "ユーザー名もしくはパスワードが間違っています"})
}
