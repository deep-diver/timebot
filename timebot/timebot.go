// Package timebot is a service package providing
// helper functions neeeded for time conversions
package timebot

import (
	"strings"
	"time"
)

// ParseTime takes a string and returns time.Time in time.UTC
func ParseTime(text string) (time.Time, bool) {
	const longForm = "2006-01-02 15:04 -0700"

	text = strings.Replace(text, "PST", "-0800", 1)
	text = strings.Replace(text, "KST", "+0900", 1)

	t, err := time.Parse(longForm, text)
	if err != nil {
		return t, false
	}
	return t.UTC(), true
}
