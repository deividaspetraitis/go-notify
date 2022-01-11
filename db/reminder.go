package db

import (
	"context"
	
	"github.com/deividaspetraitis/go-notify"
)

// ReminderService represents actual reminder service implementation
type ReminderService struct {
	// db handler and other states
}

func (s *ReminderService) FindReminders(ctx context.Context, filter notify.ReminderFilter) ([]notify.Reminder, error) {
	return []notify.Reminder{
		{Title: "take a walk dog"},
		{Title: "be happy"},
	}, nil
}

func (s *ReminderService) MarkDeliveryStatus(id int, status string) error {
	return nil
}
