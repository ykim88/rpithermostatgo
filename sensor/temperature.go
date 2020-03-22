package sensor

type Temperature struct {
	Error error
	value float64
}

func (t *Temperature) Celsius() float64 {
	return t.value / 1000.0
}
