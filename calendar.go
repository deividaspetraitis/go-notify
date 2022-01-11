package notify

import "context"

type CalendarEvent struct {
	Title string
}

type CalendarEventFilter struct {
	Offset      *int
	Limit       *int
}

type CalendarService interface {
	FindCalendarEvents(ctx context.Context, filter CalendarEventFilter) ([]CalendarEvent, error)
}
