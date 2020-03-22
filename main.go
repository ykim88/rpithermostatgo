package main

import (
	"RPiThermostatGo/heat"
	"RPiThermostatGo/sensor"
	"RPiThermostatGo/storage"
	"fmt"
	"log"
	"os"
)

func main() {
	logFile := setupLog()
	defer logFile.Close()

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
			log.Println(temperature.Error)
		}

		heatProvider.Next(temperature.Celsius()).Apply()
		storage.Save(temperature)
		fmt.Println(temperature.Celsius())
	}
}

func setupLog() *os.File {
	logF, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logF.Close()
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logF)
	return logF
}
