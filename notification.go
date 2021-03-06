package notify

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

// Notification consist of ...
type Notification struct {
	Reminder      []Reminder
	CalendarEvent []CalendarEvent
}

type NotificationProvider interface {
	RunScheduler(errors chan<- error) (<-chan Notification, chan<- string)
}

type NotificationService struct {
	providers []NotificationProvider
	errors    chan error
}

func NewNotificationService(providers []NotificationProvider) *NotificationService {
	return &NotificationService{
		providers: providers,
		errors:    make(chan error),
	}
}

func (ns *NotificationService) RunNotificationSender() <-chan error {
	for _, v := range ns.providers {
		ns.runNotificationSender(v.RunScheduler(ns.errors))
	}

	return ns.errors
}

func (ns *NotificationService) runNotificationSender(notifications <-chan Notification, status chan<- string) {
	go func() {
		for {
			n := <-notifications

			log.Println("send notification", n)

			// send notification over a wire
			// process status of responses

			status <- "sent failed|success state update"

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)

			if r1.Intn(5) == 0 {
				ns.errors <- errors.New("oops error!")
			}
		}
	}()
}
