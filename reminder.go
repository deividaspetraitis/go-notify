package main

type Reminder struct {
	Title string
}

// ReminderService represents actual reminder service implementation
type ReminderService struct {
	// db handler and other states
}

func (s *ReminderService) FindReminders() ([]Reminder, error) {
	return []Reminder{
		{Title: "take a walk dog"},
		{Title: "be happy"},
	}, nil
}

func (s *ReminderService) MarkDeliveryStatus(id int, status string) error {
	return nil
}
