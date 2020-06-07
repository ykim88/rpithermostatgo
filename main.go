package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"rpithermostatgo/api"
	"rpithermostatgo/heat"
	"rpithermostatgo/sensor"
	drivers "rpithermostatgo/sensor/drivers"
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

	gateway, err := storage.NewSQLiteReadingGateway(connectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	storage, err := storage.NewSQLiteWritingGateway(connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}

	api := api.New(api.TemperatureController(gateway))

	go heatControl(storage)

	api.Up()
}

func heatControl(gateway storage.StorageWritingGateway) {

	heat := heat.NewHeat()

	driver, err := drivers.NewSysfsDriver()
	if err != nil {
		log.Fatal(err)
	}
	sensor := sensor.TemperatureSensor(driver)

	defer sensor.Close()

	temperatureChanges := sensor.AuditChanges()

	for {
		temperature := <-temperatureChanges
		if err := temperature.Error(); err != nil {
			log.Println(err)
			continue
		}

		_ = heat.NextState(temperature.Celsius())
		gateway.Save(temperature)
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
