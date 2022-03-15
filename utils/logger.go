package utils

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/kuritaeiji/gin-practice/myvalidator"
)

func SettingLogFile() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for key, value := range myvalidator.Validators {
			v.RegisterValidation(key, value)
		}
	}
}
