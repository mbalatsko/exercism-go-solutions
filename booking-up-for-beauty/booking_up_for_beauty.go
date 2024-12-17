package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	return parsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	return parsed.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"
	parsed, _ := time.Parse(layout, date)
	return parsed.Hour() >= 12 && parsed.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsed := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", parsed.Weekday(), parsed.Month(), parsed.Day(), parsed.Year(), parsed.Hour(), parsed.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, &time.Location{})
}
