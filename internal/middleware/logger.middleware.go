package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		c.Next()

		statusCode := c.Writer.Status()
		latency := time.Since(startTime)

		log.Printf("[GIN LOGGER] %v | %3d | %13v | %15s | %-7s %#v",
			startTime.Format("2006-01-02 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
