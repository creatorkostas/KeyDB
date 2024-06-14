package cmd_api

import (
	"net/http"
	"time"

	"github.com/creatorkostas/KeyDB/api"
	"github.com/creatorkostas/KeyDB/internal/middleware"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)

func Setup_router(router *gin.Engine) {

	router.Use(helmet.Default())

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(stats.RequestStats())

	router.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		// limit 10 qps/clientIp and permit bursts of at most 10 tokens, and the limiter liveness time duration is 1 hour
		return rate.NewLimiter(rate.Every(100*time.Millisecond), 10), time.Minute * 60
	}, func(c *gin.Context) {
		c.AbortWithStatus(429) // handle exceed rate limit request
	}))
}

func Add_endpointis(router *gin.Engine) {

	router.POST("/api/register", api.Register)

	authorized := router.Group("/api/:user")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/get", api.GetValue)
		authorized.GET("/get_all", api.GetAll)
		authorized.POST("/set", api.SetValues)

		// // nested group
		admin := authorized.Group("/admin")
		admin.Use(middleware.IsAdmin())
		{
			admin.GET("/save", api.Save)
			admin.GET("/load", api.Load)
			admin.GET("/stats", func(c *gin.Context) {
				c.JSON(http.StatusOK, stats.Report())
			})

		}
	}
}
