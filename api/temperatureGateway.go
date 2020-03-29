package api

type TemperatureGateway interface {
	GetLast() (float64, error)
}
