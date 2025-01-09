package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := os.Getenv("ADMIN_SECRET")


		// TODO: taking header authorization
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(401,gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// TODO: validate header with admin password
		if auth != key {
			c.JSON(401,gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// TODO: continue request to handler
		c.Next()
	}
}