package drivers

import (
	"github.com/ykim88/go-onewire/netlink/connector/w1"
	"github.com/ykim88/go-onewire/netlink/connector/w1/ds18b20"
)

//not use many error
func NewOnewireDriver() (*onewireDriver, error) {

	w1Conn, err := w1.Dial()
	if err != nil {
		return nil, err
	}

	masters, err := w1Conn.ListMasters()
	if err != nil {
		return nil, err
	}
	devices := make([]ds18b20.Device, 0)
	for _, master := range masters {
		slaves, err := w1Conn.ListSlaves(master)
		if err != nil {
			return nil, err
		}
		for _, slave := range slaves {
			devices = append(devices, ds18b20.MakeDevice(w1Conn, slave))
		}
	}
	driver := onewireDriver{conn: w1Conn, device: devices[0]}
	return &driver, nil
}

type onewireDriver struct {
	conn   *w1.Conn
	device ds18b20.Device
}

func (d *onewireDriver) Read() (float64, error) {

	scratchpad, err := d.device.Read()
	if err != nil {
		return -56, err
	}
	err = d.device.ConvertT()
	if err != nil {
		return -56, err
	}

	return float64(scratchpad.Temperature.Float32()), nil
}

func (d *onewireDriver) Close() error {

	return d.conn.Close()
}
