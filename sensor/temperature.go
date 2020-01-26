package sensor

type Temperature interface {
	Celsius() (float64, error)
}

type temperature struct {
	value float64
}

func (t *temperature) Celsius() (float64, error) {
	return t.value / 1000.0, nil
}
