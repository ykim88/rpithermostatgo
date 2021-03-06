package api

import (
	"errors"
	"log"
	"rpithermostatgo/heat/sensor"

	"github.com/google/uuid"
)

type EventBus interface {
	Push(temperature sensor.Temperature)
	get() chan sensor.Temperature
}

func NewEventBus() EventBus {
	return &eventBus{bus: make(chan sensor.Temperature)}
}

type eventBus struct {
	bus chan sensor.Temperature
}

func (b *eventBus) Push(temperature sensor.Temperature) {
	b.bus <- temperature
}

func (b *eventBus) get() chan sensor.Temperature {
	return b.bus
}

type Subscription struct {
	id    string
	event chan sensor.Temperature
}

func (s *Subscription) Event() <-chan sensor.Temperature {
	return s.event
}

func newDispatcer(bus EventBus) *clientHandler {
	limit := 4
	return &clientHandler{subscribers: make(map[string]chan sensor.Temperature, limit), pool: newPool(limit), subscriptionRequests: make(chan chan<- *subscriptionRequest), unsubscriptionRequests: make(chan string), source: bus}

}

type clientHandler struct {
	subscribers            map[string]chan sensor.Temperature
	pool                   *channelPool
	source                 EventBus
	subscriptionRequests   chan chan<- *subscriptionRequest
	unsubscriptionRequests chan string
}

type subscriptionRequest struct {
	subscription *Subscription
	error        error
}

func (d *clientHandler) subscribe() *Subscription {
	request := make(chan *subscriptionRequest)

	d.subscriptionRequests <- request

	responce := <-request
	if responce.error != nil {
		log.Println(responce.error.Error())
		return nil
	}
	return responce.subscription
}

func (d *clientHandler) unsubscribe(subscription *Subscription) {
	d.unsubscriptionRequests <- (*subscription).id
	(*subscription).event = nil
}

func (d *clientHandler) start() {
	go d.requestHandler()
}

func (d *clientHandler) requestHandler() {
	for {
		select {
		case temperature := <-d.source.get():
			for _, subscriber := range d.subscribers {
				subscriber <- temperature
			}
		case request := <-d.subscriptionRequests:
			id := uuid.NewString()
			channel, err := d.pool.Get()
			if err != nil {
				request <- &subscriptionRequest{subscription: nil, error: err}
				continue
			}

			d.subscribers[id] = channel
			sb := Subscription{id: id, event: channel}
			request <- &subscriptionRequest{subscription: &sb, error: nil}
		case uuid := <-d.unsubscriptionRequests:
			if channel, exists := d.subscribers[uuid]; exists {
				delete(d.subscribers, uuid)
				d.pool.Put(channel)
			}

		}
	}
}

type channelPool struct {
	length int
	pool   chan chan sensor.Temperature
}

func newPool(length int) *channelPool {
	pool := make(chan chan sensor.Temperature, length)

	for i := 0; i < length; i++ {
		pool <- make(chan sensor.Temperature)
	}

	return &channelPool{length: length, pool: pool}
}

func (c *channelPool) Get() (chan sensor.Temperature, error) {
	if c.length > 0 {
		c.length--
		return <-c.pool, nil
	}
	return nil, errors.New("Other requests not accepted")
}

func (c *channelPool) Put(channel chan sensor.Temperature) {
	c.length++
	c.pool <- channel
}
