package routers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// IamTeapot     godoc
// @Summary      Teapot endpoint
// @Description  Responds with "I am a teapot!"
// @Tags         teapot
// @Produce      html
// @Success      418
// @Router       /teapot [get]
func IamTeapot(ctx *fiber.Ctx) (err error) {
	err = ctx.Status(http.StatusTeapot).SendString("I am a teapot")
	return
}
