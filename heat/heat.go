package heat

type heat interface {
	nextState(temperature float64) error
}

func newHeat() heat {
	return &simple_Heat{stateProvider: newHeatStateProvider()}
}

type simple_Heat struct {
	stateProvider heatStateProvider
}

func (h *simple_Heat) nextState(temperature float64) error {

	return h.stateProvider.getState(temperature).apply()
}
