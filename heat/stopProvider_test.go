package heat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStop(t *testing.T) {
	maxTemperature := float64(23)
	provider := stopProvider{maxTemperature: maxTemperature, stop: new(stopStateFake)}

	handler := provider.Next(maxTemperature + 0.01)

	assert.IsType(t, &stopStateFake{}, handler)
}

func TestLimitMax(t *testing.T) {
	maxTemperature := float64(23)
	provider := stopProvider{maxTemperature: maxTemperature, stop: new(stopStateFake)}

	handler := provider.Next(float64(maxTemperature))

	assert.Nil(t, handler)
}

func TestNotStop(t *testing.T) {
	maxTemperature := float64(23)
	provider := stopProvider{maxTemperature: maxTemperature, stop: new(stopStateFake)}

	handler := provider.Next(maxTemperature - 0.001)

	assert.Nil(t, handler)
}

type stopStateFake struct {
}

func (h *stopStateFake) Execute() error {

	return nil
}
