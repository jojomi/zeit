package zeit

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Time denotes a time within a day.
type Time struct {
	Hours   int
	Minutes int
}

// NewTime creates a new Time by given hours and minutes.
func NewTime(hours, minutes int) (Time, error) {
	if hours < 0 || hours > 23 {
		return Time{}, fmt.Errorf("invalid hours: %d", hours)
	}
	if minutes < 0 || minutes > 59 {
		return Time{}, fmt.Errorf("invalid minutes: %d", minutes)
	}
	return Time{
		Hours:   hours,
		Minutes: minutes,
	}, nil
}

// NewTimeFromTime creates a new Time by given time.Time.
func NewTimeFromTime(input time.Time) Time {
	return MustNewTimeParsed(input.Format("15:04"))
}

// NewTimeParsed creates a new Time from a string.
func NewTimeParsed(input string) (Time, error) {
	t := Time{}
	err := t.Parse(input)
	return t, err
}

// MustNewTimeParsed creates a new Time from a string, panicking if that does not work.
func MustNewTimeParsed(input string) Time {
	t := Time{}
	err := t.Parse(input)
	if err != nil {
		panic(err)
	}
	return t
}

// MustNewTimeParsedPointer creates a pointer to a new Time from a string, panicking if that does not work.
func MustNewTimeParsedPointer(input string) *Time {
	result := MustNewTimeParsed(input)
	return &result
}

// Parse sets the time to the interpretation of the given string.
func (t *Time) Parse(input string) error {
	zeit := regexp.MustCompile(`^(\d?\d):?(\d{2})$`) // loose to enable better error messages below
	data := zeit.FindStringSubmatch(input)

	// no match?
	if len(data) < 2 {
		return fmt.Errorf("invalid time format: %s", input)
	}

	// validate hours
	hours := data[1]
	h, _ := strconv.Atoi(hours) // thanks to the regexp this must be parseable
	if h >= 24 || h < 0 {
		return fmt.Errorf("invalid hours in %s", input)
	}

	// minutes
	minutes := data[2]
	m, _ := strconv.Atoi(minutes) // thanks to the regexp this must be parseable
	if m >= 60 || m < 0 {
		return fmt.Errorf("invalid minutes in %s", input)
	}

	// only modify state if all checks have passed!
	t.Hours = h
	t.Minutes = m

	return nil
}

// MinutesSinceMidnight returns the number of minutes that have passed to this Time since midnight.
func (t *Time) MinutesSinceMidnight() int {
	return t.Hours*60 + t.Minutes
}

// Before returns true if this Time is before another time.
func (t *Time) Before(o *Time) bool {
	return t.Hours < o.Hours || (t.Hours == o.Hours && t.Minutes < o.Minutes)
}

// BeforeEqual returns true if this Time is before another time or they are the same.
func (t *Time) BeforeEqual(o *Time) bool {
	return t.Hours < o.Hours || (t.Hours == o.Hours && t.Minutes <= o.Minutes)
}

// After returns true if this Time is after another time.
func (t *Time) After(o *Time) bool {
	return t.Hours > o.Hours || (t.Hours == o.Hours && t.Minutes > o.Minutes)
}

// AfterEqual returns true if this Time is after another time or they are the same.
func (t *Time) AfterEqual(o *Time) bool {
	return t.Hours > o.Hours || (t.Hours == o.Hours && t.Minutes >= o.Minutes)
}

// Diff returns the absolute Duration between this Time and another Time.
func (t *Time) Diff(o *Time) Duration {
	if t.Before(o) {
		return NewDurationFromMinutes(o.MinutesSinceMidnight() - t.MinutesSinceMidnight())
	}
	return NewDurationFromMinutes(t.MinutesSinceMidnight() - o.MinutesSinceMidnight())
}

// AsClock returns this Time as a string like 9:25.
func (t *Time) AsClock() string {
	return fmt.Sprintf("%d:%02d", t.Hours, t.Minutes)
}

// AsLongClock returns this Time as a string like 09:25.
func (t *Time) AsLongClock() string {
	return fmt.Sprintf("%02d:%02d", t.Hours, t.Minutes)
}

// String returns this Time as a readable string.
func (t *Time) String() string {
	return t.AsClock()
}

func IsValidTime(input string) bool {
	_, err := NewTimeParsed(input)
	return err == nil
}
