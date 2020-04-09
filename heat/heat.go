package heat

type Heat interface {
	NextState(temperature float64) error
}

type heat struct {
	stateProvider heatStateProvider
}

func (h *heat) NextState(temperature float64) error {

	return h.stateProvider.GetState(temperature).Apply()
}

func NewHeat() Heat {

	return &heat{stateProvider: newHeatStateProvider()}
}
