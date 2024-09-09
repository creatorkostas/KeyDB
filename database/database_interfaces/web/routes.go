package web_interface

import (
	"net/http"

	api "github.com/creatorkostas/KeyDB/database/database_api"
	web_api "github.com/creatorkostas/KeyDB/database/database_interfaces/web/api"
	middleware "github.com/creatorkostas/KeyDB/database/database_interfaces/web/middleware"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
)

func SetVariables() {
	api.RouterSetupFunc = Setup_router
	api.RouterAddEndpointsFunc = Add_endpoints
}

func Setup_router(router *gin.Engine) {

	router.Use(helmet.Default())

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(stats.RequestStats())
	// router.Use(middleware.Cors())

	// router.Use(middleware.AddLimiter())
}

func Add_endpoints(router *gin.Engine) {
	// http.Handle("/", http.FileServer(http.Dir("./frontend/build")))
	// http.Handle("/", http.FileServer(http.Dir("public/")))
	// router.LoadHTMLGlob("frontend/build/*.html")
	// router.LoadHTMLFiles("frontend/build/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/api/v1/register", web_api.Register)

	authorized := router.Group("/api/v1/:user")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/get", middleware.CanGet(), web_api.GetValue)
		authorized.GET("/get_all", middleware.CanGet(), web_api.GetValue)
		authorized.GET("/set", middleware.CanGetAdd(), web_api.SetValues)
		authorized.GET("/stats", middleware.CanGetAnalytics(), web_api.GetStats)

		// // nested group
		admin := authorized.Group("/admin")
		admin.Use(middleware.IsAdmin())
		{
			admin.GET("/save", web_api.Save)
			admin.GET("/load", web_api.Load)
			admin.GET("/disableAdmin", web_api.DisableAdmin)
			admin.GET("/enableAdmin", web_api.EnableAdmin)
		}
	}
}
