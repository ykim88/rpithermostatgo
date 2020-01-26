package sensor

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"RPiThermostatGo/sensor/filesystem"
)

const sensorsPath = "/sys/bus/w1/devices/%s/hwmon/hwmon1/temp1_input"

type TemperatureSensor interface {
	AuditChanges() (<-chan Temperature, error)
	Close()
}

func Sensor() (TemperatureSensor, error) {
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
	fileWatcher filesystem.FsWatcher
}

func (t *temperaturSensor) Close() {
	t.running = false
	t.fileWatcher.Stop()
}

func (s *temperaturSensor) AuditChanges() (<-chan Temperature, error) {

	s.fileWatcher = filesystem.NewFsWatcher(s.fullPath)
	changeEvent, err := s.fileWatcher.Start()
	if err != nil {
		return nil, err
	}

	ch := make(chan Temperature)

	go func(sensor *temperaturSensor, valueRead chan Temperature) {
		defer close(valueRead)
		sensor.running = true

		for sensor.running {
			ev, ok := <-changeEvent
			if !ok {
				sensor.running = false
				continue
			}
			if ev.Error != nil {
				continue
			}
			valueRead <- readSensor(sensor.fullPath)
		}

	}(s, ch)
	return ch, nil
}

func readSensor(fullpath string) Temperature {
	data, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return &invalidTemperature{error: err}
	}

	raw := string(data)

	temperatureValue, err := strconv.ParseFloat(raw[:len(raw)-1], 64)
	if err != nil {
		return &invalidTemperature{error: err}
	}

	return &temperature{value: temperatureValue}
}
