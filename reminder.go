package notify

import (
	"context"
	"time"
)

type Reminder struct {
	Title string
}

type ReminderFilter struct {
	BeforeStart *time.Duration
	Language    []string
	Offset      *int
	Limit       *int
}

type ReminderService interface {
	FindReminders(ctx context.Context, filter ReminderFilter) ([]Reminder, error)
	MarkDeliveryStatus(id int, status string) error
}
