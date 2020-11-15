package storage

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	cache := temperatureCache{}

	cache.update(10.2)

	assert.NotNil(t, cache.current)
	assert.Equal(t, 10.2, *cache.current)
}

func TestGet(t *testing.T) {
	cache := temperatureCache{}

	_, err := cache.getLast()

	assert.NotNil(t, err)
	assert.Equal(t, "Cache are Empty", err.Error())
}

func TestGetAfterUpdate(t *testing.T) {
	cache := temperatureCache{}
	cache.update(10.2)

	current, err := cache.getLast()

	assert.Nil(t, err)
	assert.NotNil(t, current)
	assert.Equal(t, 10.2, current)
}

func TestUno(t *testing.T) {
	cache := temperatureCache{}
	cache.update(10.5)

	w := sync.WaitGroup{}
	w.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			c, _ := cache.getLast()
			fmt.Println(c)
		}
		w.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			cache.update(float64(i * 10))
		}
		w.Done()
	}()
	w.Wait()
}
