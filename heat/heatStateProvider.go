package heat

type HeatStateProvider interface {
	Next(temperature float64) HeatState
}

func NewHeatStateProvider() HeatStateProvider {

	passThrough := &heatStateChain{heatStateProvider: new(passThroughProvider), next: nil}
	start := &heatStateChain{heatStateProvider: newStartStateProvider(18), next: passThrough}
	return &heatStateChain{heatStateProvider: newStopStateProvider(23), next: start}
}

type heatStateChain struct {
	heatStateProvider HeatStateProvider
	next              *heatStateChain
}

func (h *heatStateChain) Next(temperature float64) HeatState {

	handler := h.heatStateProvider.Next(temperature)
	if handler != nil {
		return handler
	}

	return h.next.Next(temperature)
}
