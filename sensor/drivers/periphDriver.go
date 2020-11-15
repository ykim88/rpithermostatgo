package drivers

import (
	"rpithermostatgo/sensor"

	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/devices/ds18b20"
	"periph.io/x/periph/experimental/host/netlink"
	"periph.io/x/periph/host"
)

//not use many error
func NewPeriphDriver() (sensor.SensorDriver, error) {
	host.Init()

	oneWBus, err := netlink.New(001)
	if err != nil {
		return nil, err
	}

	addres, err := oneWBus.Search(false)
	if err != nil {
		return nil, err
	}

	sensor, err := ds18b20.New(oneWBus, addres[0], 11)
	if err != nil {
		return nil, err
	}

	return &periphDriver{sensor: sensor, bus: oneWBus /*resolution: 12*/, env: new(physic.Env)}, nil
}

type periphDriver struct {
	sensor *ds18b20.Dev
	bus    *netlink.OneWire
	// resolution int
	env *physic.Env
}

func (d *periphDriver) Read() (float64, error) {

	err := d.sensor.Sense(d.env)
	// err := ds18b20.ConvertAll(d.bus, d.resolution)
	if err != nil {
		return -56, err
	}

	// temperature, err := d.sensor.LastTemp()
	// if err != nil {
	// 	return -56, err
	// }
	return d.env.Temperature.Celsius(), nil
}

func (d *periphDriver) Close() error {

	return d.bus.Close()
}
