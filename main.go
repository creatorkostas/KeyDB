package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/creatorkostas/KeyDB/api"
	"github.com/creatorkostas/KeyDB/internal/middleware"
	"github.com/creatorkostas/KeyDB/internal/tools"
	"github.com/gin-gonic/gin"
)

func cleanup() {
	tools.Save()
}

var accounts gin.Accounts = gin.Accounts{"kostas": "1"}

// var accounts gin.Accounts = make(gin.Accounts, 5)

// accounts["kostas"] = "1"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(gin.BasicAuth(accounts))
	// Per route middleware, you can add as many as you desire.
	// router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	authorized := router.Group("/api")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/get", api.GetValue)
		authorized.GET("/set", api.SetValues)

		// // nested group
		admin := authorized.Group("/admin")
		admin.Use(middleware.IsAdmin())
		{
			// admin.GET("/getSample", api.GetSample)
			admin.GET("/save", api.Save)
			admin.GET("/load", api.Load)

		}
	}

	tools.Load()
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		cleanup()
		<-c
		os.Exit(1)
	}()

}
