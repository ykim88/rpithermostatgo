package storage

import (
	"RPiThermostatGo/sensor"
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //sqliteDriver
)

type StorageWritingGateway interface {
	Save(celsius sensor.Temperature) error
}

func NewSQLiteWritingGateway(connectionString string) (StorageWritingGateway, error) {

	connection, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, err
	}
	return &sqliteWritingGateway{connection: connection}, nil
}

type sqliteWritingGateway struct {
	connection *sql.DB
}

func (g *sqliteWritingGateway) Save(temperature sensor.Temperature) error {

	insert, err := g.connection.Prepare("INSERT INTO Temperature (Value) VALUES (?)")
	defer insert.Close()
	if err != nil {
		return err
	}
	celsius := temperature.Celsius()
	_, err = insert.Exec(celsius)
	if err != nil {
		return err
	}
	return nil
}
