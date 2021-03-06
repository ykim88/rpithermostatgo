package sensor

import (
	"time"
)

type Sensor interface {
	AuditChanges() <-chan Temperature
	Close()
}

type SensorDriver interface {
	Read() (float64, error)
	Close() error
}

func TemperatureSensor(driver SensorDriver) Sensor {

	return &temperaturSensor{driver: driver}
}

type temperaturSensor struct {
	driver  SensorDriver
	running bool
}

func (s *temperaturSensor) Close() {
	s.running = false
}

func (s *temperaturSensor) AuditChanges() <-chan Temperature {
	ch := make(chan Temperature)

	go func(sensor *temperaturSensor, valueRead chan Temperature) {
		defer close(valueRead)
		defer sensor.driver.Close()

		sensor.running = true

		for sensor.running {
			value, err := sensor.driver.Read()
			valueRead <- &temperature{_error: err, value: value}
			<-time.After(time.Second * 5)
		}

	}(s, ch)
	return ch
}
