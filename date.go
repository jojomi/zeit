package zeit

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/now"
)

// Date is a single day with no relevant time of day.
type Date struct {
	value time.Time
}

// NewDateFromTime creates a new Date from a standard library time.Time.
func NewDateFromTime(input time.Time) Date {
	return Date{
		value: input,
	}
}

// NewDateToday creates a new Date from today/now.
func NewDateToday() Date {
	return NewDateFromTime(time.Now())
}

// Begin returns the begin of this day (midnight).
func (x Date) Begin() time.Time {
	return now.With(x.value).BeginningOfDay()
}

// End returns the end of this day (one moment before midnight).
func (x Date) End() time.Time {
	return now.With(x.value).EndOfDay()
}

// Noon returns noon of this day.
func (x Date) Noon() time.Time {
	return x.End().Add(-12 * time.Hour).Add(1 * time.Nanosecond)
}

// Equals returns true if and only if two Date structs are denoting the same day.
func (x Date) Equals(o Date) bool {
	layout := "20060102"
	return x.End().Format(layout) == (o.End().Format(layout))
}

func (x Date) String() string {
	return x.End().Format("2006-01-02")
}

func (x Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

func (x Date) UnmarshalJSON(data []byte) error {
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return err
	}

	newDate, err := ParseDate(strVal)
	if err != nil {
		return err
	}

	x.value = newDate.value
	return nil
}
