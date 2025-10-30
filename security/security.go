package security

import (
	"os"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
