package main

import (
	"RPiThermostatGo/api"
	"RPiThermostatGo/heat"
	"RPiThermostatGo/sensor"
	"RPiThermostatGo/storage"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	logFile := setupLog()
	defer logFile.Close()

	user, err := user.Current()
	if err != nil {
		log.Fatal(err.Error())
	}

	connectionString := fmt.Sprintf("%s/storage/RPiThermostatGo.db", user.HomeDir)
	err = storage.CreateDbSchemaIfNotExists(connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	gateway, err := storage.NewSQLiteReadingGateway(connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	api := api.New(api.TemperatureController(gateway))

	go heatControl(connectionString)

	api.Up()
}

func heatControl(connectionString string) {
	storage, err := storage.NewSQLiteWritingGateway(connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}
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
