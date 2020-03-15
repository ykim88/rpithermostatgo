package handler

type Handler interface {
}

type HandlerProvider interface {
	Get(temperature float64) Handler
}

type handlerBaseProvider struct {
	handlerProvider HandlerProvider
	next            *handlerBaseProvider
}

func (h *handlerBaseProvider) Get(temperature float64) Handler {

	handler := h.handlerProvider.Get(temperature)
	if handler != nil {
		return handler
	}

	return h.next.Get(temperature)
}

type passThroughProvider struct{}

func (p *passThroughProvider) Get(temperature float64) Handler {
	return new(passThroughHandler)
}

type passThroughHandler struct{}
