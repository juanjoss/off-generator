package main

import (
	"github.com/juanjoss/off-generator/event"

	"github.com/jasonlvhit/gocron"
)

func main() {
	// registering events
	eventProvider := event.NewEventProvider()

	eventProvider.RegisterEvent(&event.ProductOrder{})
	eventProvider.RegisterEvent(&event.ProductAddedToSSD{})
	eventProvider.RegisterEvent(&event.UserRegistration{})

	/**
	launch scheduler to generate random registered events
	*/
	gocron.Every(1).Minute().Do(eventProvider.SelectEvent)

	<-gocron.Start()
}
