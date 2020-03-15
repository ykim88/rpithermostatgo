package storage_test

import (
	"RPiThermostatGo/storage"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"
)

var ConnectionString string
var Database *TestDatabase

func TestMain(m *testing.M) {
	setup()
	pass := m.Run()
	tearDown()
	os.Exit(pass)
}

func setup() error {
	ConnectionString = fmt.Sprintf("/tmp/test%s.db", time.Now().Format(time.RFC3339Nano))
	storage.CreateDbSchemaIfNotExists(ConnectionString)
	Database = new(TestDatabase)
	return nil
}

func tearDown() {
	os.Remove(ConnectionString)
}

type TestDatabase struct {
}

func (db *TestDatabase) ReadAll() []float64 {
	connection, _ := sql.Open("sqlite3", ConnectionString)
	query, _ := connection.Prepare("SELECT Value FROM Temperature")
	rows, err := query.Query()
	if err != nil {
		return nil
	}
	savedtemperature := make([]float64, 0)
	for rows.Next() {
		var tmp float64
		rows.Scan(&tmp)
		savedtemperature = append(savedtemperature, tmp)
	}
	return savedtemperature
}
