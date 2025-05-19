package logging

import "github.com/gin-gonic/gin"

type Responder interface {
	Error(ctx *gin.Context, status int, msg string)
}

var ResponseJSON Responder

type JSONResponder struct{}

func InitJSONResponder() *JSONResponder {
	return &JSONResponder{}
}

func (r *JSONResponder) Error(ctx *gin.Context, status int, msg string) {
	ctx.JSON(status, gin.H{
		"success": false,
		"msg":     msg,
	})
}
