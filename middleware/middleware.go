package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	apiKey := c.GetHeader("x-api-key")
	if apiKey != "secret-key" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid API key",
		})
		return
	}
	c.Next()
}
