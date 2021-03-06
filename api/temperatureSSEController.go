package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TemperatureSSEController(bus EventBus) *sseController {
	dispatcher := newDispatcer(bus)
	dispatcher.start()
	return &sseController{clientHandler: dispatcher}
}

type sseController struct {
	clientHandler *clientHandler
}

func (controller *sseController) bind(router *gin.Engine) {
	router.GET("temperature/realTime", controller.realTime)
}

func (controller *sseController) realTime(c *gin.Context) {
	w := c.Writer
	flusher, ok := w.(http.Flusher)
	if !ok {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Streaming unsupported"))
		return
	}

	subscription := controller.clientHandler.subscribe()
	if subscription == nil {
		c.AbortWithError(http.StatusTooManyRequests, fmt.Errorf("unableToSubscribe"))
	}
	defer controller.clientHandler.unsubscribe(subscription)
	close := w.(http.CloseNotifier).CloseNotify()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	done := false
	for !done {
		select {
		case temperature, open := <-subscription.Event():
			if !open {
				break
			}
			c.SSEvent("message", temperature.Celsius())
			flusher.Flush()

		case <-close:
			done = true
		}
	}
	c.AbortWithStatus(http.StatusOK)
}
