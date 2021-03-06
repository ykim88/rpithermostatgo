package heat

type passThroughProvider struct{}

func (p *passThroughProvider) getState(temperature float64) heatState {
	return new(passThroughState)
}
