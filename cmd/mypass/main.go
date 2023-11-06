package main

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	_ "github.com/sourcehaven/mypass-godbridge/pkg/ini"
)

// @title         MyPass API
// @version       0.1.0
// @license.name  MIT
// @host          localhost:7277
// @BasePath      /api
func main() {
	mypass := app.New()
	mypass.Start()
}
