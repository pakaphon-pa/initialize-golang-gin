package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (a *Adapter) Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")

}
