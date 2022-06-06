package event

import (
	"log"
	"math/rand"
)

type Event interface {
	Handle()
	Type() string
}

type EventProvider struct {
	events []Event
}

func NewEventProvider() *EventProvider {
	return &EventProvider{
		events: []Event{},
	}
}

func (ep *EventProvider) RegisterEvent(e Event) {
	ep.events = append(ep.events, e)
}

func (ep *EventProvider) SelectEvent() {
	index := rand.Intn(len(ep.events))
	event := ep.events[index]
	log.Println(event.Type())
	event.Handle()
}
