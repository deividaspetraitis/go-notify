package db

import (
	"context"

	"github.com/deividaspetraitis/go-notify"
)

// CalendarService represents actual calendar service implementation
type CalendarService struct {
	// db handler and other states
}

func (s *CalendarService) FindCalendarEvents(ctx context.Context, filter notify.CalendarEventFilter) ([]notify.CalendarEvent, error) {
	return []notify.CalendarEvent{
		{Title: "meeting with manager at 10am"},
		{Title: "take car to service"},
	}, nil
}
