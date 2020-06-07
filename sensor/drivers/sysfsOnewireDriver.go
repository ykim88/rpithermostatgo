package drivers

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func NewSysfsDriver() (*sysfsDriver, error) {

	raw, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	if err != nil {
		return nil, err
	}

	sensorsID := strings.Split(string(raw), "\n")
	if len(sensorsID) > 0 {
		sensorsID = sensorsID[:len(sensorsID)-1]
	}

	sensors := make([]*sysfsDriver, len(sensorsID))
	for i, sensorID := range sensorsID {
		sensors[i] = &sysfsDriver{filePath: "/sys/bus/w1/devices/" + sensorID + "/w1_slave"}
	}

	return sensors[0], nil
}

type sysfsDriver struct {
	filePath string
}

func (driver *sysfsDriver) Close() error {
	return nil
}

func (driver *sysfsDriver) Read() (float64, error) {

	data, err := ioutil.ReadFile(driver.filePath)
	if err != nil {
		return -56, err
	}

	raw := string(data)

	err = crcCheck(raw)
	if err != nil {
		return -56, err
	}

	return parseTemperature(raw)
}

func crcCheck(raw string) error {

	if !strings.Contains(raw, " YES") {
		return errors.New("CRC check failed")
	}

	return nil
}

func parseTemperature(raw string) (float64, error) {

	i := strings.LastIndex(raw, "t=")
	if i == -1 {
		return -56, errors.New("failed to read sensor temperature")
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
	if err != nil {
		return -56, errors.New("failed parsing")
	}
	return c / 1000, nil
}
