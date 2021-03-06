package heat

type stopState struct{}

func (s *stopState) apply() error {
	return nil
}
