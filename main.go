package main

import (
	"os"
	"os/signal"
	"syscall"

	cmd_api "github.com/creatorkostas/KeyDB/cmd/api"
	"github.com/creatorkostas/KeyDB/internal"
	"github.com/creatorkostas/KeyDB/internal/handlers"
	"github.com/creatorkostas/KeyDB/internal/tools"
	"github.com/gin-gonic/gin"
)

// var DB_filename string = "db.gob"

func cleanup() {
	tools.SaveDB(internal.DB_filename, &handlers.DB)
	tools.SaveDB(internal.Accounts_filename, &handlers.Accounts)
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	// Creates a router without any middleware by default
	router := gin.New()

	cmd_api.Setup_router(router)
	cmd_api.Add_endpointis(router)

	tools.LoadDB(internal.DB_filename, &handlers.DB)
	tools.LoadDB(internal.Accounts_filename, &handlers.Accounts)

	router.Run(":8080")

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		cleanup()
		<-c
		os.Exit(1)
	}()

}
