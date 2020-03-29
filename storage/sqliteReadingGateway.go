package storage

import (
	"RPiThermostatGo/api"
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //sqliteDriver
)

func NewSQLiteReadingGateway(connectionString string) (api.TemperatureGateway, error) {

	return &sqliteReadingGateway{connectionString: connectionString}, nil
}

type sqliteReadingGateway struct {
	connectionString string
}

func (g *sqliteReadingGateway) GetLast() (float64, error) {
	connection, err := open(g.connectionString)
	defer connection.Close()
	if err != nil {
		return -56, err
	}

	var temperature float64
	err = connection.QueryRow("SELECT Value FROM Temperature ORDER BY DateTime DESC LIMIT 1").Scan(&temperature)
	if err != nil && sql.ErrNoRows != err {
		return -56, err
	}
	return temperature, nil
}

func open(connectionString string) (*sql.DB, error) {
	return sql.Open("sqlite3", connectionString)
}
