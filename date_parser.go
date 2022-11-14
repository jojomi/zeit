package zeit

import (
	"fmt"
	"strconv"
	"time"
)

// ParseDate parses a string as a date.
// Supported formats:
//   - YYYY-MM-DD
//   - today, yesterday, tomorrow
//   - -2, +4 (days from now)
func ParseDate(input string) (Date, error) {
	switch input {
	case "yesterday":
		return parseWithDayDiff(-1), nil
	case "today":
		return parseWithDayDiff(0), nil
	case "tomorrow":
		return parseWithDayDiff(1), nil
	}

	intVal, err := strconv.Atoi(input)
	if err == nil {
		return parseWithDayDiff(intVal), nil
	}

	res, err := time.Parse("2006-01-02", input)
	if err == nil {
		return NewDateFromTime(res), nil
	}

	return Date{}, fmt.Errorf("could not parse date: '%s'", input)
}

func parseWithDayDiff(diffDays int) Date {
	n := time.Now()
	return NewDateFromTime(n.AddDate(0, 0, diffDays))
}
