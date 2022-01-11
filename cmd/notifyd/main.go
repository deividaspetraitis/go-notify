package main

import (
	"log"

	"github.com/deividaspetraitis/go-notify"
	"github.com/deividaspetraitis/go-notify/db"
)

func main() {
	// TODO: context

	// construct notification service by injecting notification data providers
	notificationservice := notify.NewNotificationService([]notify.NotificationProvider{
			notify.NewReminderNotificationProvider(&db.ReminderService{}),
			notify.NewCalendarNotificationProvider(&db.CalendarService{}),
	})

	errors := notificationservice.RunNotificationSender()

	// separate routine watching and reporting errors
	go func() {
		if err := <-errors; err != nil {
			log.Println("received an error", err)
		}
	}()

	// TODO: ctr+c signal handling
	for {
	}
}
