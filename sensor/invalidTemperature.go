package sensor

type invalidTemperature struct {
	error error
}

func (t *invalidTemperature) Celsius() (float64, error) {
	return 0.0, t.error
}
