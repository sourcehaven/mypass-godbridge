package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IamTeapot     godoc
// @Summary      Teapot endpoint
// @Description  Responds with "I am a teapot!"
// @Tags         teapot
// @Produce      html
// @Success      418
// @Router       /teapot [get]
func IamTeapot(c *gin.Context) {
	c.JSON(http.StatusTeapot, "I am a teapot!")
}
