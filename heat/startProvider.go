package heat

func newStartStateProvider(minTemperature float64) heatStateProvider {

	return &startProvider{minTemperature: minTemperature, start: new(startState)}
}

type startProvider struct {
	minTemperature float64
	start          HeatState
}

func (p *startProvider) GetState(temperature float64) HeatState {

	if temperature < p.minTemperature {

		return p.start
	}

	return nil
}
