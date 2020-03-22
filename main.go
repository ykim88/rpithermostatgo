package main

import (
	"RPiThermostatGo/heat"
	"RPiThermostatGo/sensor"
	"RPiThermostatGo/storage"
	"fmt"
	"log"
)

func main() {

	connectionString := "/tmp/RPiThermostatGo.db"
	storage.CreateDbSchemaIfNotExists(connectionString)
	storage := storage.NewSQLiteStorageGateway(connectionString)
	heatProvider := heat.NewHeatStateProvider()

	sensor, err := sensor.TemperatureSensor()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sensor.Close()

	temperatureChanges := sensor.AuditChanges()

	for {
		temperature := <-temperatureChanges
		if temperature.Error != nil {
			fmt.Println(temperature.Error.Error())
		}
		heatProvider.Next(temperature.Celsius()).Execute()
		storage.Save(temperature)
		fmt.Println(temperature.Celsius())
	}
}
