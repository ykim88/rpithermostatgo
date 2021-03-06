package heat

func newStopStateProvider(maxTemperature float64) heatStateProvider {

	return &stopProvider{maxTemperature: maxTemperature, stop: new(stopState)}
}

type stopProvider struct {
	maxTemperature float64
	stop           heatState
}

func (p *stopProvider) getState(temperature float64) heatState {

	if temperature > p.maxTemperature {

		return p.stop
	}

	return nil
}
