package storage_test

import (
	"RPiThermostatGo/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	gateway := storage.NewSQLiteStorageGateway(ConnectionString)
	temperature := new(validTemperature)
	expected, _ := temperature.Celsius()

	err := gateway.Save(temperature)

	assert.Nil(t, err)
	savedTemperature := Database.ReadAll()
	assert.Equal(t, expected, savedTemperature[0])
}

type validTemperature struct {
}

func (t *validTemperature) Celsius() (float64, error) {
	return 22.0, nil
}
