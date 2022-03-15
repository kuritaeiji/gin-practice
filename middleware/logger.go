package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var buf []byte
		if ctx.Request.Body != nil {
			buf, _ = ioutil.ReadAll(ctx.Request.Body)
		}
		r1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		r2 := ioutil.NopCloser(bytes.NewBuffer(buf))

		buf2 := new(bytes.Buffer)
		buf2.ReadFrom(r1)
		fmt.Println(buf2.String())

		ctx.Request.Body = r2
	}
}
