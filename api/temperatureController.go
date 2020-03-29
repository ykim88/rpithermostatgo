package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TemperatureController(gateway TemperatureGateway) *temperatreController {
	return &temperatreController{temperatureGateway: gateway}
}

type temperatreController struct {
	temperatureGateway TemperatureGateway
}

func (controller *temperatreController) bind(router *gin.Engine) {
	router.GET("temperature/Current", controller.current)
}

func (controller *temperatreController) current(c *gin.Context) {
	temperature, err := controller.temperatureGateway.GetLast()
	if err != nil {

		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "")
	}
	c.String(http.StatusOK, "%f", temperature)
}
