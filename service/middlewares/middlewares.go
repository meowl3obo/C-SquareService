package middlewares

import (
	"fmt"
	"net/http"

	provider "C-SquaredService/provider"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if provider.ConnErr != nil {
			c.String(http.StatusInternalServerError, "connect db error")
			c.Abort()
			return
		}
		c.Next()
	}
}
