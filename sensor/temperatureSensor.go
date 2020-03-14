package sensor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const sensorsPath = "/sys/bus/w1/devices/%s/w1_slave"

type Sensor interface {
	AuditChanges() <-chan Temperature
	Close()
}

func TemperatureSensor() (Sensor, error) {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	if err != nil {
		return nil, err
	}

	sensors := strings.Split(string(data), "\n")
	if len(sensors) > 0 {
		sensors = sensors[:len(sensors)-1]
	}

	return &temperaturSensor{fullPath: fmt.Sprintf(sensorsPath, sensors[0])}, nil
}

type temperaturSensor struct {
	fullPath    string
	running     bool
	timer       *time.Timer
}

func (s *temperaturSensor) Close() {
	s.running = false
	s.timer.Stop()
}

func (s *temperaturSensor) AuditChanges() <-chan Temperature {

	s.timer = time.NewTimer(time.Second * 5)
	ch := make(chan Temperature)

	go func(sensor *temperaturSensor, valueRead chan Temperature) {
		defer close(valueRead)
		sensor.running = true

		for sensor.running {
			valueRead <- readSensor(sensor.fullPath)
			<-s.timer.C
		}

	}(s, ch)
	return ch
}

func readSensor(fullpath string) Temperature {
	data, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return &invalidTemperature{error: err}
	}

	raw := string(data)

	indexTemp := strings.LastIndex(raw, "t=")
	if indexTemp == -1 {
		return &invalidTemperature{error: errors.New("failed to read sensor temperature")}
	}

	temperatureValue, err := strconv.ParseFloat(raw[indexTemp+2:len(raw)-1], 64)
	if err != nil {
		return &invalidTemperature{error: err}
	}

	return &temperature{value: temperatureValue}
}
