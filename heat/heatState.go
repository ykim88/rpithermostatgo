package heat

type heatState interface {
	apply() error
}
