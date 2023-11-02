package main

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	_ "github.com/sourcehaven/mypass-godbridge/pkg/ini"
	"github.com/sourcehaven/mypass-godbridge/pkg/security/jwt"
)

// @title         MyPass API
// @version       0.1.0
// @license.name  MIT
// @host          localhost:7277
// @BasePath      /api
func main() {
	token, err := jwt.CreateAccessToken("mypass", jwt.TokenOptions{Fresh: true})
	if err != nil {
		return
	}
	validatedClaims, err := jwt.ValidateFreshAccessToken(token)
	println(validatedClaims)
	mypass := app.New()
	mypass.Start()
}
