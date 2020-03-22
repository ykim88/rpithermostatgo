package heat

type HeatState interface {
	Apply() error
}
