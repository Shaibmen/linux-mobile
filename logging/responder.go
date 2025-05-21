package logging

import (
	"server/models"

	"github.com/gin-gonic/gin"
)

type Responder interface {
	Error(ctx *gin.Context, status int, msg string, err error)
}

var ResponseJSON Responder

type JSONResponder struct{}

func InitJSONResponder() *JSONResponder {
	return &JSONResponder{}
}

func (r *JSONResponder) Error(ctx *gin.Context, status int, msg string, err error) {
	ctx.JSON(status, models.HttpResponse{
		Out:   msg,
		Error: err,
	})
}
