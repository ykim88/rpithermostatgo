package sensor

type Temperature interface {
	Error() error
	Celsius() float64
}

type temperature struct {
	_error error
	value  float64
}

func (t *temperature) Error() error {
	return t._error
}

func (t *temperature) Celsius() float64 {
	return t.value
}
