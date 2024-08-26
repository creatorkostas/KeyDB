package web_api

import (
	"net/http"

	"github.com/creatorkostas/KeyDB/database/database_api/web/api"
	middleware "github.com/creatorkostas/KeyDB/database/database_api/web/middleware"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
)

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
	router.GET("/api/v1/register", api.Register)

	authorized := router.Group("/api/v1/:user")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/get", middleware.CanGet(), api.GetValue)
		authorized.GET("/get_all", middleware.CanGet(), api.GetValue)
		authorized.POST("/set", middleware.CanGetAdd(), api.SetValues)
		authorized.GET("/stats", middleware.CanGetAnalytics(), api.GetStats)

		// // nested group
		admin := authorized.Group("/admin")
		admin.Use(middleware.IsAdmin())
		{
			admin.GET("/save", api.Save)
			admin.GET("/load", api.Load)
			admin.GET("/disableAdmin", api.DisableAdmin)
			admin.GET("/enableAdmin", api.EnableAdmin)
		}
	}
}
