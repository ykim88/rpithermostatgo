package heat

type HeatState interface {
	Execute() error
}
