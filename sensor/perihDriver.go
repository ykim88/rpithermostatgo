package sensor

import (
	"periph.io/x/periph/devices/ds18b20"
	"periph.io/x/periph/experimental/host/netlink"
	"periph.io/x/periph/host"
)

func NewPeriphDriver() (*periphDriver, error) {
	host.Init()

	oneWBus, err := netlink.New(001)
	if err != nil {
		return nil, err
	}

	addres, err := oneWBus.Search(false)
	if err != nil {
		return nil, err
	}

	sensor, err := ds18b20.New(oneWBus, addres[0], 10)
	if err != nil {
		return nil, err
	}

	return &periphDriver{sensor: sensor, bus: oneWBus, resolution: 10}, nil
}

type periphDriver struct {
	sensor     *ds18b20.Dev
	bus        *netlink.OneWire
	resolution int
}

func (d *periphDriver) Read() (float64, error) {

	err := ds18b20.ConvertAll(d.bus, d.resolution)
	if err != nil {
		return -56, err
	}

	temperature, err := d.sensor.LastTemp()
	if err != nil {
		return -56, err
	}

	return temperature.Celsius(), nil
}

func (d *periphDriver) Close() error {

	return d.bus.Close()
}
