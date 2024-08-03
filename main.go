package main

import (
	"flag"

	cmd_api "github.com/creatorkostas/KeyDB/database/database_api/cmd"
	web_api "github.com/creatorkostas/KeyDB/database/database_api/web"
	database "github.com/creatorkostas/KeyDB/database/database_core"
	internal "github.com/creatorkostas/KeyDB/database/database_core/conf"

	// "github.com/creatorkostas/KeyDB/internal/database/database_test"
	"github.com/creatorkostas/KeyDB/database/database_core/persistance"
	"github.com/creatorkostas/KeyDB/database/database_core/users"
	"github.com/gin-gonic/gin"
)

// var DB_filename string = "db.gob"

var router = gin.New()

func cleanup() {
	persistance.SaveToFile(internal.DB_filename, &database.DB)
	persistance.SaveToFile(internal.Accounts_filename, &users.Accounts)
	persistance.Operations <- "||exit||"
}

func initialize() {
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	persistance.LoadFromFile(internal.DB_filename, &database.DB)
	persistance.LoadFromFile(internal.Accounts_filename, &users.Accounts)

	// c := make(chan os.Signal, 2)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// go func() {
	// 	<-c
	// 	cleanup()
	// 	os.Exit(1)
	// }()

	persistance.Start_writers(1)

	web_api.Setup_router(router)
	web_api.Add_endpoints(router)

}

func main() {

	devMode := false
	conf_path := "config.yaml"
	flag.BoolVar(&devMode, "dev", devMode, "enable dev mode")
	flag.StringVar(&conf_path, "conf", conf_path, "Set the config file")
	flag.Parse()

	internal.Load_configs(conf_path)
	initialize()
	cmd_api.Cmd_start()
	router.Run(":8080")
}
