package heat

type stopState struct{}

func (s *stopState) Execute() error {

	return nil
}
