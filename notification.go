package notify

import "log"

// Notification consist of ...
type Notification struct {
	Reminder []Reminder
	CalendarEvent []CalendarEvent
}

type NotificationProvider interface {
	RunScheduler(errors chan<- error) (<-chan Notification, chan<- string)
}

type NotificationService struct {
	DataProviders []NotificationProvider
}

func (ns *NotificationService) RunNotificationSender(notifications <-chan Notification, status chan<- string) {
	go func() {
		for {
			n := <-notifications

			log.Println("send notification", n)

			// send notification over a wire
			// process status of responses

			status <- "sent failed|success state update"
		}
	}()
}
