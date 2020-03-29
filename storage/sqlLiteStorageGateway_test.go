package storage_test

import (
	"RPiThermostatGo/sensor"
	"RPiThermostatGo/storage"
	"testing"

	_ "github.com/mattn/go-sqlite3" //sqliteDriver
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	gateway := storage.NewSQLiteStorageGateway(ConnectionString)
	temperature := &sensor.Temperature{}
	expected := temperature.Celsius()

	err := gateway.Save(temperature)

	assert.Nil(t, err)
	savedTemperature := Database.ReadAll()
	assert.Equal(t, expected, savedTemperature[0])
}
