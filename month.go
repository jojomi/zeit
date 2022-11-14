package zeit

import (
	"fmt"
	"github.com/jinzhu/now"
	"math"
	"time"
)

// Month is a month in a year. Supports navigating back and forth as well as printing in different styles.
type Month struct {
	month int
	year  int
}

func NewMonth(month, year int) (Month, error) {
	if month < 1 || month > 12 {
		return Month{}, fmt.Errorf("invalid month value (must be >= 1 and <= 12): %d", month)
	}
	return Month{
		month: month,
		year:  year,
	}, nil
}

func MustNewMonth(month, year int) Month {
	m, err := NewMonth(month, year)
	if err != nil {
		panic(err)
	}
	return m
}

func NewMonthFromTime(input time.Time) Month {
	month, _ := NewMonth(int(input.Month()), input.Year())
	return month
}

func (x Month) Month() time.Month {
	return time.Month(x.month)
}

func (x Month) MonthNumber() int {
	return x.month
}

func (x Month) Year() int {
	return x.year
}

func (x Month) WithAddedMonths(monthCount int) Month {
	m := x.Year()*12 + x.MonthNumber() - 1 + monthCount
	newYear := int(math.Floor(float64(m) / 12))
	newMonth := (m % 12) + 1
	return MustNewMonth(newMonth, newYear)
}

func (x Month) Equals(other Month) bool {
	return x.Year() == other.Year() && x.Month() == other.Month()
}

func (x Month) Before(other Month) bool {
	if x.Year() < other.Year() {
		return true
	}
	return x.Year() == other.Year() && x.Month() < other.Month()
}

func (x Month) After(other Month) bool {
	if x.Year() > other.Year() {
		return true
	}
	return x.Year() == other.Year() && x.Month() > other.Month()
}

func (x Month) IsCurrent(referenceDate time.Time) bool {
	return x.Equals(NewMonthFromTime(referenceDate))
}

func (x Month) String() string {
	return fmt.Sprintf("%d/%d", x.month, x.year)
}

func (x Month) HumanString() string {
	monthString := time.Month(x.month).String()
	return fmt.Sprintf("%s %d", monthString, x.year)
}

func (x Month) Beginning(loc *time.Location) time.Time {
	return now.With(time.Date(x.Year(), x.Month(), 1, 12, 0, 0, 0, loc)).BeginningOfMonth()
}

func (x Month) End(loc *time.Location) time.Time {
	return now.With(time.Date(x.Year(), x.Month(), 1, 12, 0, 0, 0, loc)).EndOfMonth()
}

// MonthsUntil returns all months between this one and the given (later) Month, including both of them.
func (x Month) MonthsUntil(other Month) []Month {
	months := make([]Month, 0)

	// reverse order?
	if x.After(other) {
		return months
	}

	var (
		counter = 0
	)
	for {
		currentMonth := x.WithAddedMonths(counter)
		months = append(months, currentMonth)
		if currentMonth.Equals(other) {
			break
		}
		counter++
	}

	return months
}
