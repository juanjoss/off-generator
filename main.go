package main

import (
	"time"

	"github.com/juanjoss/off-generator/event"

	"github.com/go-co-op/gocron"
)

func main() {
	/*
		registering events
	*/
	eventProvider := event.NewEventProvider()

	eventProvider.RegisterEvent(&event.ProductOrder{})
	eventProvider.RegisterEvent(&event.ProductAddedToSSD{})
	eventProvider.RegisterEvent(&event.UserRegistration{})

	/*
		launch scheduler to generate random registered events
	*/
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Minute().Do(eventProvider.SelectEvent)

	s.StartBlocking()
}
