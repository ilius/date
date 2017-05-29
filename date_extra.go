package date

import "time"

// NewPtr creates a new *Date
func NewPtr(year int, month time.Month, day int) *Date {
	d := New(year, month, day)
	return &d
}
