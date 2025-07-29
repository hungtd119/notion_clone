package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recover middleware để handle panic
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			c.Abort()
		}
	}()
	c.Next()
}

// Cors middleware để handle CORS
func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}

// Logger middleware để log requests
func Logger(c *gin.Context) {
	// Log request info
	c.Next()

	// Log response info
	status := c.Writer.Status()
	if status >= 400 {
		// Log error responses
	}
}
