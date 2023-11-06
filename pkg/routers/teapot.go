package routers

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/app"
	"net/http"
)

// IamTeapot     godoc
// @Summary      Teapot endpoint
// @Description  Responds with "I am a teapot!"
// @Tags         teapot
// @Produce      html
// @Success      418
// @Router       /teapot [get]
func IamTeapot(ctx app.Ctx) (err error) {
	ctx.GetLogger().Debug("This is hell!")
	err = ctx.Status(http.StatusTeapot).SendString("I am a teapot")
	return
}
