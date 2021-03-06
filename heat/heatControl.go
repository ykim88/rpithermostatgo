package heat

import (
	"log"
	"rpithermostatgo/heat/sensor"
)

type HeatControl interface {
	Run()
}

func NewHeatControl(driver sensor.SensorDriver, listeners ...func(sensor.Temperature)) HeatControl {

	sensor := sensor.TemperatureSensor(driver)
	return &heatControl{sensor: sensor, heat: newHeat(), listeners: listeners}
}

type heatControl struct {
	heat      heat
	sensor    sensor.Sensor
	listeners []func(sensor.Temperature)
}

func (c *heatControl) Run() {
	go func() {
		temperatureChanges := c.sensor.AuditChanges()
		defer c.sensor.Close()

		for {
			temperature := <-temperatureChanges
			if err := temperature.Error(); err != nil {
				log.Println(err)
				continue
			}
			c.notifyTemperatureChange(temperature)
			_ = c.heat.nextState(temperature.Celsius())
		}
	}()
}

func (c *heatControl) notifyTemperatureChange(t sensor.Temperature) {

	for _, listener := range c.listeners {
		listener(t)
	}
}
