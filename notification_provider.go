package main

import (
	"log"
	"time"
)

type ReminderNotificationProvider struct {
	service ReminderService
	sleep   time.Duration
	out     chan Notification
	in      chan string
}

func NewReminderNotificationProvider() *ReminderNotificationProvider {
	return &ReminderNotificationProvider{
		service: ReminderService{},
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
			reminders, err := p.service.FindReminders()
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
	sleep time.Duration
	out   chan Notification
	in    chan string
}

func NewCalendarNotificationProvider() *CalendarNotificationProvider {
	return &CalendarNotificationProvider{
		sleep: 3 * time.Second,
		out:   make(chan Notification),
		in:    make(chan string),
	}
}

func (p *CalendarNotificationProvider) RunScheduler(errors chan<- error) (<-chan Notification, chan<- string) {
	// routine for data fetch scheduler
	go func() {
		for {
			p.out <- Notification{Calendar: []string{"Meeting with manager 10 AM"}}
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
