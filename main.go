package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	cmd_api "github.com/creatorkostas/KeyDB/cmd/api"
	"github.com/creatorkostas/KeyDB/internal"
	"github.com/creatorkostas/KeyDB/internal/database"
	aof "github.com/creatorkostas/KeyDB/internal/persistance"
	"github.com/creatorkostas/KeyDB/internal/tools"
	"github.com/creatorkostas/KeyDB/internal/users"
	"github.com/gin-gonic/gin"
)

// var DB_filename string = "db.gob"
var router = gin.New()

func cleanup() {
	tools.SaveToFile(internal.DB_filename, &database.DB)
	tools.SaveToFile(internal.Accounts_filename, &users.Accounts)
	aof.Operations <- "||exit||"
}

func initialize() {
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	tools.LoadFromFile(internal.DB_filename, &database.DB)
	tools.LoadFromFile(internal.Accounts_filename, &users.Accounts)

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		cleanup()
		<-c
		os.Exit(1)
	}()

	aof.Start_writers(1)

	cmd_api.Setup_router(router)
	cmd_api.Add_endpointis(router)

}

func main() {

	devMode := false
	flag.BoolVar(&devMode, "dev", devMode, "enable dev mode")
	flag.Parse()

	// fmt.Println(database_test.Run_write_test(100))
	// fmt.Println(database_test.Run_read_test(10000))

	initialize()
	router.Run(":8080")

}
