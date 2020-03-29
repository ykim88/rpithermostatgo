package storage

import (
	"RPiThermostatGo/api"
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //sqliteDriver
)

func NewSQLiteReadingGateway(connectionString string) (api.TemperatureGateway, error) {
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, err
	}
	return &sqliteReadingGateway{connection: db}, nil
}

type sqliteReadingGateway struct {
	connection *sql.DB
}

func (g *sqliteReadingGateway) GetLast() (float64, error) {

	var temperature float64
	err := g.connection.QueryRow("SELECT Value FROM Temperature ORDER BY DateTime DESC LIMIT 1").Scan(&temperature)
	if err != nil && sql.ErrNoRows != err {
		return -56, err
	}
	return temperature, nil
}
