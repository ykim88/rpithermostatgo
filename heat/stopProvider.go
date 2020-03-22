package heat

func newStopStateProvider(maxTemperature float64) HeatStateProvider {

	return &stopProvider{maxTemperature: maxTemperature, stop: new(stopState)}
}

type stopProvider struct {
	maxTemperature float64
	stop           HeatState
}

func (p *stopProvider) Next(temperature float64) HeatState {

	if temperature > p.maxTemperature {

		return p.stop
	}

	return nil
}