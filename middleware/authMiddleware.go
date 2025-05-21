package middleware

import (
	"net/http"
	"server/environment"
	"server/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.HttpResponse{Out: "не авторизован"})
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		if token != environment.Env.AccessKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.HttpResponse{Out: "не правильный токен"})
			return
		}

		c.Next()
	}
}
