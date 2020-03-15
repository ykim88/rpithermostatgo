package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStop(t *testing.T) {
	maxTemperature := float64(23)
	provider := stopProvider{maxTemperature: maxTemperature, stopHandler: new(stopHandlerFake)}

	handler := provider.Get(maxTemperature + 0.01)

	assert.IsType(t, &stopHandlerFake{}, handler)
}

func TestLimitMax(t *testing.T) {
	maxTemperature := float64(23)
	provider := stopProvider{maxTemperature: maxTemperature, stopHandler: new(stopHandlerFake)}

	handler := provider.Get(float64(maxTemperature))

	assert.Nil(t, handler)
}

func TestNotStop(t *testing.T) {
	maxTemperature := float64(23)
	provider := stopProvider{maxTemperature: maxTemperature, stopHandler: new(stopHandlerFake)}

	handler := provider.Get(maxTemperature - 0.001)

	assert.Nil(t, handler)
}

type stopHandlerFake struct {
}
