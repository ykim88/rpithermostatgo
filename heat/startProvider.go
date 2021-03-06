package heat

func newStartStateProvider(minTemperature float64) heatStateProvider {

	return &startProvider{minTemperature: minTemperature, start: new(startState)}
}

type startProvider struct {
	minTemperature float64
	start          heatState
}

func (p *startProvider) getState(temperature float64) heatState {

	if temperature < p.minTemperature {

		return p.start
	}

	return nil
}
