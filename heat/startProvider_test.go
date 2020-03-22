package heat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	minTemperature := float64(18)
	provider := startProvider{minTemperature: minTemperature, start: new(startStateFake)}

	handler := provider.Next(minTemperature - 0.001)

	assert.IsType(t, &startStateFake{}, handler)
}

func TestLimitMin(t *testing.T) {
	minTemperature := float64(18)
	provider := startProvider{minTemperature: minTemperature, start: new(startStateFake)}

	handler := provider.Next(minTemperature)

	assert.Nil(t, handler)
}

func TestNotStart(t *testing.T) {
	minTemperature := float64(18)
	provider := startProvider{minTemperature: minTemperature, start: new(startStateFake)}

	handler := provider.Next(minTemperature + 0.001)

	assert.Nil(t, handler)
}

type startStateFake struct {
}

func (h *startStateFake) Apply() error {

	return nil
}
