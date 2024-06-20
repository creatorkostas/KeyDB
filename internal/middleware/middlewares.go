package middleware

import (
	"net/http"
	"time"

	"github.com/creatorkostas/KeyDB/internal/api"
	"github.com/creatorkostas/KeyDB/internal/users"
	"github.com/gin-gonic/gin"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)

func getIP(c *gin.Context) string {
	return c.ClientIP() // limit rate by client ip
}

func limiter(c *gin.Context) (*rate.Limiter, time.Duration) {
	var acc = c.MustGet("Account").(*users.Account)
	// limit 10 qps/clientIp and permit bursts of at most 10 tokens, and the limiter liveness time duration is 1 hour
	return rate.NewLimiter(rate.Every(acc.Burst_time), acc.Burst_tokens), acc.Rate_reset
}

func abort(c *gin.Context) {
	c.AbortWithStatus(429) // handle exceed rate limit request
}

func AddLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit.NewRateLimiter(getIP, limiter, abort)
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		// t := time.Now()
		var api_key, api_found = c.GetQuery("api_key")
		// var user, user_found = c.GetQuery("user")

		var user = c.Param("user")

		if !api_found {
			c.JSON(http.StatusUnauthorized, api.JsonResponce{Message: "api_key req"})
			c.Abort()
		}

		var acc = users.Get_account(user)

		if acc == nil {
			c.JSON(http.StatusUnauthorized, api.JsonResponce{Message: "account not found"})
			c.Abort()
		} else {
			if acc.Tokens == 0 {
				c.JSON(429, api.JsonResponce{Message: "Your tokens have reach zero!"})
				c.Abort()
			} else {
				acc.Tokens -= 1
			}
			c.Set("Account", acc)

			if acc.Api_key == api_key && acc.Username == user {
				c.Set("pass", true)
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, api.JsonResponce{Message: "wrong user - api"})
				c.Abort()
			}
		}
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {

		var acc = c.MustGet("Account").(*users.Account)

		if acc.IsAdmin() {
			c.Next()
			c.Set("Admin", true)
		} else {
			c.Set("Admin", false)
			c.Abort()
		}

	}
}

func CanGetAnalytics() gin.HandlerFunc {
	return func(c *gin.Context) {

		var acc = c.MustGet("Account").(*users.Account)

		if acc.CanGetAnalytics() {
			c.Next()
		} else {
			c.Abort()
		}

	}
}

func CanGet() gin.HandlerFunc {
	return func(c *gin.Context) {

		var acc = c.MustGet("Account").(*users.Account)

		if acc.CanGet() {
			c.Next()
		} else {
			c.Abort()
		}

	}
}

func CanGetAdd() gin.HandlerFunc {
	return func(c *gin.Context) {

		var acc = c.MustGet("Account").(*users.Account)

		if acc.CanAdd() {
			c.Next()
		} else {
			c.Abort()
		}

	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
