package heat

type passThroughProvider struct{}

func (p *passThroughProvider) GetState(temperature float64) HeatState {
	return new(passThroughState)
}
