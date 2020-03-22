package storage

import (
	"RPiThermostatGo/sensor"
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //sqliteDriver
)

type StorageGateway interface {
	Save(celsius *sensor.Temperature) error
}

func NewSQLiteStorageGateway(connectionString string) StorageGateway {

	return &sqlLiteStorageGateway{connectionString: connectionString}
}

type sqlLiteStorageGateway struct {
	connectionString string
}

func (g *sqlLiteStorageGateway) Save(temperature *sensor.Temperature) error {
	connection, err := g.openConnection()
	if err != nil {
		return err
	}
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
	return nil
}

func (g *sqlLiteStorageGateway) openConnection() (connection *sql.DB, err error) {
	connection, err = sql.Open("sqlite3", g.connectionString)
	if err != nil {
		return nil, err
	}
	return connection, err
}
