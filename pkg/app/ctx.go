package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Ctx struct {
	fiber.Ctx
}

type CtxInterface interface {
	App() *App
	GetLogger() *logrus.Logger
}

func (ctx *Ctx) App() *App {
	return ctx.App()
}

func (ctx *Ctx) GetLogger() *logrus.Logger {
	app := ctx.App()
	return app.Logger
}
