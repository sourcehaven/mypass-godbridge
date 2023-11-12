package main

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	_ "github.com/sourcehaven/mypass-godbridge/docs"
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"github.com/sourcehaven/mypass-godbridge/pkg/ini"
	_ "github.com/sourcehaven/mypass-godbridge/pkg/ini"
	"github.com/sourcehaven/mypass-godbridge/pkg/routers"
)

func initRouters(app *fiber.App) {
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}))
	api := app.Group("/api")
	{
		api.Get("/teapot", routers.IamTeapot)

		auth := api.Group("/auth")
		{
			auth.Post("/register", routers.RegisterUser)
			auth.Post("/activate/:token", routers.ActivateUser)
			auth.Post("/login", routers.LoginUser)
		}
	}
}

// @title         MyPass API
// @version       0.1.0
// @license.name  MIT
// @host          localhost:7277
// @BasePath      /api
func main() {
	srv := app.New(app.Cfg)
	logrus.SetLevel(app.Cfg.LogLevel)
	ini.InitApp()
	initRouters(srv.App)
	srv.StartServerWithGracefulShutdown()
}
