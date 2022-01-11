package main

import "log"

func main() {
	errors := make(chan error)

	// construct notification service by injecting notification data providers
	notificationservice := NotificationService{
		DataProviders: []NotificationProvider{
			NewReminderNotificationProvider(),
			NewCalendarNotificationProvider(),
		},
	}

	// run notification sender
	for _, v := range notificationservice.DataProviders {
		notificationservice.RunNotificationSender(v.RunScheduler(errors))
	}

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
