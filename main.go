package main

import (
	"RPiThermostatGo/sensor"
	"fmt"
	"log"
)

func main() {
	sensor, err := sensor.Sensor()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sensor.Close()

	temperatureChanges, err := sensor.AuditChanges()
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		temperature := <-temperatureChanges
		value, err := temperature.Celsius()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(value)
	}
}
