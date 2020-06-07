package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func TemperatureController(gateway TemperatureGateway) *temperatreController {
	return &temperatreController{temperatureGateway: gateway, indexHandler: http.FileServer(http.Dir("./web"))}
}

type temperatreController struct {
	indexHandler       http.Handler
	temperatureGateway TemperatureGateway
}

func (controller *temperatreController) bind(router *gin.Engine) {
	dir, _ := os.Getwd()
	fmt.Println(dir)

	router.GET("/", controller.index)
	router.Static("/home", "./web")
	router.GET("temperature/current", controller.current)
}

func (controller *temperatreController) index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/home")
}

func (controller *temperatreController) current(c *gin.Context) {
	temperature, err := controller.temperatureGateway.GetLast()
	if err != nil {

		log.Println(err.Error())
		c.String(http.StatusInternalServerError, "")
	}
	c.String(http.StatusOK, "%f", temperature)
}
