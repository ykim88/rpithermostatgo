package sensor

import (
	"RPiThermostatGo/sensor/driver"
	"time"
)

const sensorsPath = "/sys/bus/w1/devices/%s/w1_slave"

type Sensor interface {
	AuditChanges() <-chan Temperature
	Close()
}

type sensorDriver interface {
	Read() (float64, error)
	Close() error
}

func TemperatureSensor() (Sensor, error) {
	// data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	// if err != nil {
	// 	return nil, err
	// }

	// sensors := strings.Split(string(data), "\n")
	// if len(sensors) > 0 {
	// 	sensors = sensors[:len(sensors)-1]
	// }
	driver, err := driver.NewPeriphDriver()
	if err != nil {
		return nil, err
	}
	return &temperaturSensor{driver: driver}, nil
}

type temperaturSensor struct {
	driver  sensorDriver
	running bool
	timer   *time.Timer
}

func (s *temperaturSensor) Close() {
	s.running = false
	s.timer.Stop()
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
