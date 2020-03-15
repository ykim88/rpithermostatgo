package handler

// func TestPass(t *testing.T) {

// 	provider := handlerProvider{maxTemperature: 23, minTemperature: 18, startHandler: &startHandlerFake{}, stopHandler: new(stopHandlerFake), passHandler: new(passHandlerFake)}

// 	handler := provider.Get(float64(22.99))

// 	assert.IsType(t, &passHandlerFake{}, handler)
// }

type passHandlerFake struct {
}
