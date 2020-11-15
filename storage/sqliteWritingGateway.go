package storage

import (
	"rpithermostatgo/sensor"

	_ "github.com/mattn/go-sqlite3" //sqliteDriver
)

type StorageWritingGateway interface {
	Save(celsius sensor.Temperature) error
}

func NewSQLiteWritingGateway(connectionString string) (StorageWritingGateway, error) {

	return &sqliteWritingGateway{connectionString: connectionString, cache: cache}, nil
}

type sqliteWritingGateway struct {
	cache            temperatureCache
	connectionString string
}

func (g *sqliteWritingGateway) Save(temperature sensor.Temperature) error {
	connection, err := open(g.connectionString)
	if err != nil {
		return err
	}
	defer connection.Close()

	insert, err := connection.Prepare("INSERT INTO Temperature (Value) VALUES (?)")
	defer insert.Close()
	if err != nil {
		return err
	}
	celsius := temperature.Celsius()
	_, err = insert.Exec(celsius)
	if err != nil {
		return err
	}
	g.cache.update(celsius)
	return nil
}
