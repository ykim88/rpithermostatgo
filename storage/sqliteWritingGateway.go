package storage

import (
	"log"
	"rpithermostatgo/heat/sensor"

	_ "github.com/mattn/go-sqlite3" //sqliteDriver
)

type StorageWritingGateway interface {
	Save(celsius sensor.Temperature)
}

func NewSQLiteWritingGateway(connectionString string) StorageWritingGateway {

	return &sqliteWritingGateway{connectionString: connectionString, cache: cache}
}

type sqliteWritingGateway struct {
	cache            *temperatureCache
	connectionString string
}

func (g *sqliteWritingGateway) Save(temperature sensor.Temperature) {
	connection, err := open(g.connectionString)
	if err != nil {
		log.Println(err.Error())
	}
	defer connection.Close()

	insert, err := connection.Prepare("INSERT INTO Temperature (Value) VALUES (?)")
	defer insert.Close()
	if err != nil {
		log.Println(err.Error())
	}
	celsius := temperature.Celsius()
	_, err = insert.Exec(celsius)
	if err != nil {
		log.Println(err.Error())
	}
	g.cache.update(celsius)
}
