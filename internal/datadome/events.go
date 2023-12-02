package datadome

import (
	"math/rand"
)

func (t *DataDomePayload) GetEvents() {
	events := new(eventCounters)
	events.Mousemove = rand.Intn(200)
	events.Click = rand.Intn(6)
	events.Scroll = rand.Intn(10)
	events.Touchstart = 0
	events.Touchend = 0
	events.Touchmove = 0
	events.Keydown = 0
	events.Keyup = 0
	t.events = *events
}

type eventCounters struct {
	Mousemove  int `json:"mousemove"`
	Click      int `json:"click"`
	Scroll     int `json:"scroll"`
	Touchstart int `json:"touchstart"`
	Touchend   int `json:"touchend"`
	Touchmove  int `json:"touchmove"`
	Keydown    int `json:"keydown"`
	Keyup      int `json:"keyup"`
}
