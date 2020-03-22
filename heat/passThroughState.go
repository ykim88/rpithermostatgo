package heat

type passThroughState struct{}

func (h *passThroughState) Execute() error {
	return nil
}
