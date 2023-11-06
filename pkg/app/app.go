package app

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/sourcehaven/mypass-godbridge/pkg/routers"
)

type App struct {
	*fiber.App
	Logger *logrus.Logger
	Config *Config
}

func initRouters(app *fiber.App) *fiber.App {
	app.Get("/docs/*", swagger.HandlerDefault)
	api := app.Group("/api")
	{
		api.Get("/teapot", routers.IamTeapot)
	}
	return app
}

func (app *App) Start() {
	host := app.Config.Host
	port := app.Config.Port
	addr := host + ":" + port
	_ = app.Listen(addr)
}

func New() *App {
	fib := fiber.New()
	log := logrus.New()
	cfg := NewConfig()
	fib = initRouters(fib)
	app := &App{fib, log, cfg}
	return app
}
