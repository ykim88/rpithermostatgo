package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	validFilePath := "./testFile/validFile"
	driver := fsdriver{sysfspath: validFilePath}

	value, err := driver.Read()

	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, float64(22000), value)
}

func TestNotValid(t *testing.T) {
	validFilePath := "./testFile/notValidFile"
	driver := fsdriver{sysfspath: validFilePath}

	value, err := driver.Read()

	assert.Equal(t, float64(-56000), value)
	assert.NotNil(t, err)
}
