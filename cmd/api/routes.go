package cmd_api

import (
	"github.com/creatorkostas/KeyDB/internal/api"
	"github.com/creatorkostas/KeyDB/internal/middleware"
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

	// router.Use(middleware.AddLimiter())
}

func Add_endpointis(router *gin.Engine) {

	router.POST("/api/register", api.Register)

	authorized := router.Group("/api/:user")
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
		}
	}
}
