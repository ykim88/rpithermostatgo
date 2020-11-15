package storage

import (
	"errors"
	"sync"
)

var cache temperatureCache

func init() {
	cache = temperatureCache{}
}

type temperatureCache struct {
	mutex   sync.RWMutex
	current *float64
}

func (c *temperatureCache) getLast() (float64, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.current == nil {
		return -56, errors.New("Cache are Empty")
	}

	return *c.current, nil
}

func (c *temperatureCache) update(lastTemperature float64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.current = &lastTemperature

}
