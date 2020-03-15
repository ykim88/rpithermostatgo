package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	minTemperature := float64(18)
	provider := startProvider{minTemperature: minTemperature, startHandler: new(startHandlerFake)}

	handler := provider.Get(minTemperature - 0.001)

	assert.IsType(t, &startHandlerFake{}, handler)
}

func TestLimitMin(t *testing.T) {
	minTemperature := float64(18)
	provider := startProvider{minTemperature: minTemperature, startHandler: new(startHandlerFake)}

	handler := provider.Get(minTemperature)

	assert.Nil(t, handler)
}

func TestNotStart(t *testing.T) {
	minTemperature := float64(18)
	provider := startProvider{minTemperature: minTemperature, startHandler: new(startHandlerFake)}

	handler := provider.Get(minTemperature + 0.001)

	assert.Nil(t, handler)
}

type startHandlerFake struct {
}
