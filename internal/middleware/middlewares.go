package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// // Set example variable
		if 1 == 1 {
			c.Next()
		}
		c.Set("pass", true)

		// before request

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Abort()
		c.Next()

	}
}
