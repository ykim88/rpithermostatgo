package heat

type heatStateProvider interface {
	getState(temperature float64) heatState
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

func (h *heatStateChain) getState(temperature float64) heatState {

	handler := h.heatStateProvider.getState(temperature)
	if handler != nil {
		return handler
	}

	return h.next.getState(temperature)
}
