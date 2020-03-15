package handler

type stopProvider struct {
	maxTemperature float64
	stopHandler    Handler
}

func (p *stopProvider) Get(temperature float64) Handler {

	if temperature > p.maxTemperature {

		return p.stopHandler
	}

	return nil
}
