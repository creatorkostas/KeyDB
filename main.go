package main

import (
	"flag"
	"os"
	"strconv"

	cmd_api "github.com/creatorkostas/KeyDB/database/database_api/cmd"
	"github.com/creatorkostas/KeyDB/database/database_core/conf"
	internal "github.com/creatorkostas/KeyDB/database/database_core/conf"
	// "github.com/creatorkostas/KeyDB/internal/database/database_test"
)

// var DB_filename string = "db.gob"

// var router = gin.New()

// func initialize() {
// 	// gin.SetMode(gin.ReleaseMode)
// 	gin.SetMode(gin.DebugMode)

// 	db_utils.LoadDB(internal.DB_filename)
// 	users.LoadAccounts(internal.Accounts_filename)

// 	// c := make(chan os.Signal, 2)
// 	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
// 	// go func() {
// 	// 	<-c
// 	// 	cleanup()
// 	// 	os.Exit(1)
// 	// }()

// 	persistance.Start_writers(1)

// 	web_api.Setup_router(router)
// 	web_api.Add_endpoints(router)

// }

func main() {

	cmd := false
	unix := false
	devMode := false
	conf_path := "config.yaml"
	flag.BoolVar(&devMode, "dev", devMode, "enable dev mode")
	flag.BoolVar(&cmd, "cmd", cmd, "enable cmd mode")
	flag.BoolVar(&unix, "unix", unix, "enable communication through unix port (/tmp/keydb_sock.sock)")
	flag.StringVar(&conf_path, "conf", conf_path, "Set the config file")
	flag.Parse()

	internal.Load_configs(conf_path)
	// initialize()
	var port = os.Getenv("PORT")

	conf.StartUnix = unix

	if port == "" {
		cmd_api.StartKeyDB(devMode, conf.StartWeb, strconv.Itoa(8080), conf.StartUnix)
	} else {
		cmd_api.StartKeyDB(devMode, conf.StartWeb, port, conf.StartUnix)
	}

	if cmd {
		cmd_api.Cmd_start()
	}
	// router.Run(":8080")
}
