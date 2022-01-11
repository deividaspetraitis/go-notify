package notify

import (
	"context"
	"log"
	"time"
)

type ReminderNotificationProvider struct {
	service ReminderService
	sleep   time.Duration
	out     chan Notification
	in      chan string
}

func NewReminderNotificationProvider(service ReminderService) *ReminderNotificationProvider {
	return &ReminderNotificationProvider{
		service: service,
		sleep:   5 * time.Second,
		out:     make(chan Notification),
		in:      make(chan string),
	}
}

func (p *ReminderNotificationProvider) RunScheduler(errors chan<- error) (<-chan Notification, chan<- string) {
	log.Println("running reminder notifications provider scheduler")

	// routine for data fetch scheduler
	go func() {
		for {
			reminders, err := p.service.FindReminders(context.TODO(), ReminderFilter{})
			if err != nil {
				errors <- err
				time.Sleep(p.sleep)
			}

			p.out <- Notification{Reminder: reminders}
			time.Sleep(p.sleep)
		}
	}()

	// routine for updating notification delivery status
	go func() {
		for {
			id, status := 1, <-p.in

			log.Println("Reminder notification provider received delivery status", status)

			p.service.MarkDeliveryStatus(id, status)
		}
	}()

	return p.out, p.in
}

type CalendarNotificationProvider struct {
	service CalendarService
	sleep time.Duration
	out   chan Notification
	in    chan string
}

func NewCalendarNotificationProvider(service CalendarService) *CalendarNotificationProvider {
	return &CalendarNotificationProvider{
		service: service,
		sleep: 3 * time.Second,
		out:   make(chan Notification),
		in:    make(chan string),
	}
}

func (p *CalendarNotificationProvider) RunScheduler(errors chan<- error) (<-chan Notification, chan<- string) {
	// routine for data fetch scheduler
	go func() {
		for {
			events, err := p.service.FindCalendarEvents(context.TODO(), CalendarEventFilter{})
			if err != nil {
				errors <- err
				time.Sleep(p.sleep)
			}

			p.out <- Notification{CalendarEvent: events}
			time.Sleep(p.sleep)
		}
	}()

	// routine for updating notification delivery status
	go func() {
		for {
			log.Println("Calendar notification provider received deliver status", <-p.in)
		}
	}()

	return p.out, p.in
}
