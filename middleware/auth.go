package middleware

import (
	"log"
	"net/http"
	"server/environment"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			log.Println(header)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"out": "не авторизован"})
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		if token != environment.Env.AccessKey {
			log.Println(environment.Env.AccessKey)
			log.Println(header)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"out": "не правильный токен"})
			return
		}

		c.Next()
	}
}
