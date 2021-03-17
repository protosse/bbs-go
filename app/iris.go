package app

import (
	"bbs-go/common/config"
	"bbs-go/controllers/api"
	"fmt"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func InitIris() {
	app := iris.New()
	app.Logger().SetLevel(config.Config.LogLevel)
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodHead, iris.MethodOptions},
		AllowedHeaders:   []string{"*"},
	}))

	mvc.Configure(app.Party("/api"), func(a *mvc.Application) {
		a.Party("/login").Handle(new(api.LoginController))
	})

	_ = app.Run(
		iris.Addr(fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithConfiguration(iris.Configuration{
			EnableOptimizations:     true,
			DisableInterruptHandler: true,
			TimeFormat:              "2006-01-02 15:04:05",
		}))
}
