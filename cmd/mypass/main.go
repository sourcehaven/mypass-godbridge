package main

import (
	"mypass-godbridge/pkg/app"
)

// @title         MyPass API
// @version       0.1.0
// @license.name  MIT
// @host          localhost:7277
// @BasePath      /api
func main() {
	cfg := app.Config{}
	if cfg.Env() == "development" {
		app.DummyDbInit()
	}
	mypass := app.CreateApp()
	mypass.Start()
}
