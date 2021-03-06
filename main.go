package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"rpithermostatgo/api"
	"rpithermostatgo/heat"
	"rpithermostatgo/heat/sensor/drivers"
	"rpithermostatgo/storage"
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

	gateway := storage.NewSQLiteReadingGateway(connectionString)

	controller := api.TemperatureController(gateway)
	eventbus := api.NewEventBus()
	sseController := api.TemperatureSSEController(eventbus)
	api := api.New(controller, sseController)

	driver, err := drivers.NewSysfsDriver()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewSQLiteWritingGateway(connectionString)
	heatControl := heat.NewHeatControl(driver, storage.Save, eventbus.Push)
	heatControl.Run()
	api.Up()
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
