package heat

type passThroughProvider struct{}

func (p *passThroughProvider) Next(temperature float64) HeatState {
	return new(passThroughState)
}
