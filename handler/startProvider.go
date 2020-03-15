package handler

type startProvider struct {
	minTemperature float64
	startHandler   Handler
}

func (p *startProvider) Get(temperature float64) Handler {

	if temperature < p.minTemperature {

		return p.startHandler
	}

	return nil
}
