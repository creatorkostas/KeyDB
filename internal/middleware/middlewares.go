package middleware

import (
	"net/http"

	"github.com/creatorkostas/KeyDB/api"
	"github.com/creatorkostas/KeyDB/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		// t := time.Now()
		var api_key, api_found = c.GetQuery("api_key")
		// var user, user_found = c.GetQuery("user")

		var user = c.Param("user")

		var acc = handlers.Get_account(user)

		if !api_found {
			c.IndentedJSON(http.StatusUnauthorized, api.Responce{Message: "api_key req"})
			c.Abort()
			// }
			// else if !user_found {
			// c.IndentedJSON(http.StatusUnauthorized, api.Responce{Message: "user req"})
			// c.Abort()
		} else if acc == nil {
			c.IndentedJSON(http.StatusUnauthorized, api.Responce{Message: "account not found"})
			c.Abort()
		} else {
			// if accounts[user] == api_key {
			// log.Println(acc.Api_key)
			// log.Println(acc.Username)
			// log.Println(api_key)
			// log.Println(user)
			if acc.Api_key == api_key && acc.Username == user {
				c.Set("pass", true)
				c.Next()
			} else {
				c.IndentedJSON(http.StatusUnauthorized, api.Responce{Message: "wrong user - api"})
				c.Abort()
			}
		}

		// // Set example variable

		// if c.MustGet(gin.AuthUserKey) != nil {
		// 	c.Next()
		// }
		// c.Set("pass", true)

		// before request

		// after request
		// latency := time.Since(t)
		// log.Print(latency)

		// access the status we are sending
		// status := c.Writer.Status()
		// log.Println(status)
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = c.Param("user")

		var acc = handlers.Get_account(user)

		// if accounts_roles[user] == "Admin" {
		if acc.Tier.Type == handlers.ADMIN {
			c.Next()
		} else {
			c.Abort()
		}

	}
}

// func cors() gin.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Access-Control-Allow-Origin", "*")
// 		handler.ServeHTTP(w, r)
// 	})
// }

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
