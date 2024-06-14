package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	cmd_api "github.com/creatorkostas/KeyDB/cmd/api"
	"github.com/creatorkostas/KeyDB/frontend"
	"github.com/creatorkostas/KeyDB/internal"
	"github.com/creatorkostas/KeyDB/internal/handlers"
	"github.com/creatorkostas/KeyDB/internal/middleware"
	"github.com/creatorkostas/KeyDB/internal/tools"
	"github.com/gin-gonic/gin"
)

// var DB_filename string = "db.gob"

func cleanup() {
	tools.SaveDB(internal.DB_filename, &handlers.DB)
	tools.SaveDB(internal.Accounts_filename, &handlers.Accounts)
}

func main() {

	devMode := false
	flag.BoolVar(&devMode, "dev", devMode, "enable dev mode")
	flag.Parse()

	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	// Creates a router without any middleware by default
	router := gin.New()

	// mux.Handle("/admin/", frontend.SvelteKitHandler("/admin"))

	router.GET("/admin/", gin.WrapH(frontend.SvelteKitHandler("./frontend")))
	if devMode {
		router.Use(middleware.Cors())
		fmt.Println("server running in dev mode")
	}

	cmd_api.Setup_router(router)
	cmd_api.Add_endpointis(router)
	// router.Static("/_app/immutable/", "./frontend/.svelte-kit/output/client/ ")
	router.Static("/sta", "./frontend/.svelte-kit/output/")
	// router.StaticFS("/index", http.Dir("./frontend/.svelte-kit/output/prerendered/pages/"))
	// router.StaticFS("/register", http.Dir("./frontend/.svelte-kit/output/prerendered/pages/"))

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
