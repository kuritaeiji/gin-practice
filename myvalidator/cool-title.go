package myvalidator

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var Validators = map[string]validator.Func{
	"cool": cool,
}

var cool validator.Func = func(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "Cool")
}
