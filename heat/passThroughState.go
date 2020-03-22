package heat

type passThroughState struct{}

func (h *passThroughState) Apply() error {
	return nil
}
