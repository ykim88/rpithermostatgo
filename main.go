package main

import (
	"RPiThermostatGo/sensor"
	"RPiThermostatGo/storage"
	"fmt"
	"log"
)

func main() {
	connectionString := "/tmp/RPiThermostatGo.db"
	storage.CreateDbSchemaIfNotExists(connectionString)
	storage := storage.NewSQLiteStorageGateway(connectionString)
	sensor, err := sensor.TemperatureSensor()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sensor.Close()

	temperatureChanges := sensor.AuditChanges()

	for {
		temperature := <-temperatureChanges
		value, err := temperature.Celsius()
		if err != nil {
			fmt.Println(err.Error())
		}
		storage.Save(temperature)
		fmt.Println(value)
	}
}
