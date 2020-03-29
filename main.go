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

	connectionString := "$HOME/storage/RPiThermostatGo.db"
	err := storage.CreateDbSchemaIfNotExists(connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	storage := storage.NewSQLiteStorageGateway(connectionString)
	heatProvider := heat.NewHeatStateProvider()

	sensor, err := sensor.TemperatureSensor()
	if err != nil {
		sensor.Close()
		log.Fatal(err.Error())
	}
	defer sensor.Close()

	temperatureChanges := sensor.AuditChanges()

	for {
		temperature := <-temperatureChanges

		if err := temperature.Error(); err != nil {
			log.Println(err)
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
