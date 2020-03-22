package heat

type stopState struct{}

func (s *stopState) Apply() error {

	return nil
}
