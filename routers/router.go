package routers

import (
	"net/http"

	// "github.com/casbin/casbin"
	// "github.com/getsentry/raven-go"
	// "github.com/gin-contrib/gzip"
	// "github.com/gin-contrib/sentry"
	// "github.com/gin-contrib/zap"
	// "go.uber.org/zap"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"upgrade/backend/controllers/user"
	"upgrade/backend/libs/setting"
	// "upgrade/backend/middlewares"
)

func InitRouter() *gin.Engine {
	// default engine instance
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Access", "Accept", "Authorization", "Content-Type"}
	r.Use(cors.New(config))

	// set run mode
	gin.SetMode(setting.RunMode)

	// serve docs file
	r.Static("/docs", "./docs")

	// gzip middleware
	// r.Use(gzip.Gzip(gzip.DefaultCompression))

	// serve favicon.ico
	r.Static("/favicon.ico", "/docs/img/favicon.ico")

	// serve public file
	r.StaticFS("/static", http.Dir("./public/static"))
	r.LoadHTMLFiles("public/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// middleware
	// e := casbin.NewEnforcer("conf/authz.conf", "authz_policy.csv")
	// r.Use(middlewares.Authz(e))
	// r.Use(middlewares.Logger())
	// raven.SetDSN(setting.SentryKey)
	// r.Use(sentry.Recovery(raven.DefaultClient, true))
	// logger, _ := zap.NewProduction()
	// r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// set api prefix
	api := r.Group("/api")
	api.POST("/user/register", user.AddUser)

	return r
}
