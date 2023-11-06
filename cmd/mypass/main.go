package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	_ "github.com/sourcehaven/mypass-godbridge/pkg/ini"
	"github.com/sourcehaven/mypass-godbridge/pkg/routers"
)

func initRouters(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
	api := app.Group("/api")
	{
		api.Get("/teapot", routers.IamTeapot)
	}
}

// @title         MyPass API
// @version       0.1.0
// @license.name  MIT
// @host          localhost:7277
// @BasePath      /api
func main() {
	engine := fiber.New()
	log := logrus.New()
	cfg := app.NewConfig()

	log.SetLevel(cfg.LogLevel)
	//appContext := &deps.AppContext{Config: cfg, Logger: log}

	initRouters(engine)
	addr := cfg.Host + ":" + cfg.Port
	log.Fatal(engine.Listen(addr))
}
