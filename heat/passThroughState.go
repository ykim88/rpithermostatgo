package heat

type passThroughState struct{}

func (h *passThroughState) apply() error {
	return nil
}
