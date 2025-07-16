package middleware

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		fmt.Printf("[GIN] %s - %s %s - %v\n", c.ClientIP(), c.Request.Method, c.Request.URL.Path, duration)
	}
}
