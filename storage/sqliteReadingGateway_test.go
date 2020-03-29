package storage_test

import (
	"RPiThermostatGo/storage"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	insert(22, time.Now())
	gateway, _ := storage.NewSQLiteReadingGateway(ConnectionString)

	value, err := gateway.GetLast()

	assert.Nil(t, err)
	assert.Equal(t, float64(22), value)
}

func TestReadLast(t *testing.T) {
	insert(22, time.Now().Add(time.Second))
	insert(18, time.Now())
	gateway, _ := storage.NewSQLiteReadingGateway(ConnectionString)

	value, err := gateway.GetLast()

	assert.Nil(t, err)
	assert.Equal(t, float64(22), value)
}

func insert(value float64, dateTime time.Time) {
	connection, _ := sql.Open("sqlite3", ConnectionString)
	insert, _ := connection.Prepare("INSERT INTO Temperature (Value, DateTime) VALUES (?,?)")
	defer insert.Close()
	insert.Exec(value, dateTime)
}
