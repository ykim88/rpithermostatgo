package heat

type heatStateProvider interface {
	GetState(temperature float64) HeatState
}

func newHeatStateProvider() heatStateProvider {

	passThrough := &heatStateChain{heatStateProvider: new(passThroughProvider), next: nil}
	start := &heatStateChain{heatStateProvider: newStartStateProvider(18), next: passThrough}
	return &heatStateChain{heatStateProvider: newStopStateProvider(23), next: start}
}

type heatStateChain struct {
	heatStateProvider heatStateProvider
	next              *heatStateChain
}

func (h *heatStateChain) GetState(temperature float64) HeatState {

	handler := h.heatStateProvider.GetState(temperature)
	if handler != nil {
		return handler
	}

	return h.next.GetState(temperature)
}
