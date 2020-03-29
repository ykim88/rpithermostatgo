package storage_test

import (
	"RPiThermostatGo/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	gateway, _ := storage.NewSQLiteWritingGateway(ConnectionString)
	temperature := &fakeTemperature{}
	expected := temperature.Celsius()

	err := gateway.Save(temperature)

	assert.Nil(t, err)
	savedTemperature := Database.ReadAll()
	assert.Equal(t, expected, savedTemperature[0])
}

type fakeTemperature struct {
}

func (t *fakeTemperature) Celsius() float64 {
	return 21
}

func (t *fakeTemperature) Error() error {
	return nil
}
