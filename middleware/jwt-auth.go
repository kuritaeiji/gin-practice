package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kuritaeiji/gin-practice/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEME = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEME):]

		token, err := service.NewJWTService().ValidateJWT(tokenString)
		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]:, ", claims["name"])
			log.Println("Claims[Admin]:", claims["admin"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			ctx.AbortWithStatus(401)
		}
	}
}
