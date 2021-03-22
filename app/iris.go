package app

import (
	"bbs-go/common/config"
	"bbs-go/controllers/api"
	"bbs-go/middleware"
	"bbs-go/middleware/jauth"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

type IrisServer struct {
	App    *iris.Application
	Config *config.Config
}

func NewIrisServer(config *config.Config) *IrisServer {
	server := &IrisServer{
		Config: config,
	}
	server.initIris()
	return server
}

func (s *IrisServer) initIris() {
	app := iris.New()
	app.Logger().SetLevel(s.Config.LogLevel)
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodHead, iris.MethodOptions},
		AllowedHeaders:   []string{"*"},
	}))

	jwtAccessAuth := jauth.Access()
	cacheAuth := middleware.CacheAccessAuth

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/user").Handle(new(api.LoginController))
		m.Router.Use(jwtAccessAuth, cacheAuth)
		m.Party("/user").Handle(new(api.UserController))
	})

	s.App = app
}

func (s *IrisServer) Run() {
	_ = s.App.Run(
		iris.Addr(fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithConfiguration(iris.Configuration{
			EnableOptimizations:     true,
			DisableInterruptHandler: true,
			TimeFormat:              "2006-01-02 15:04:05",
		}))
}
