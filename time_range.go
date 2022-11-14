package zeit

import (
	"encoding/json"
	"fmt"
	"github.com/juju/errors"
)

// TimeRange is a pair of start and end time while both are optional.
type TimeRange struct {
	Start *Time
	End   *Time
}

func NewTimeRange(start *Time, end *Time) TimeRange {
	return TimeRange{
		Start: start,
		End:   end,
	}
}

func (x TimeRange) Duration() (Duration, error) {
	if x.Start == nil {
		return NewDurationFromMinutes(0), errors.New("can't calculate duration: no start time set")
	}
	if x.End == nil {
		return NewDurationFromMinutes(0), errors.New("can't calculate duration: no end time set, try DurationUntilNow?")
	}
	return NewDurationFromTimes(*x.Start, *x.End), nil
}

func (x TimeRange) DurationUntilNow(now Time) Duration {
	end := x.End
	if end == nil {
		end = &now
	}

	return NewDurationFromTimes(*x.Start, *end)
}

func (x TimeRange) String() string {
	start := "*"
	if x.Start != nil {
		start = x.Start.AsClock()
	}
	end := "*"
	if x.End != nil {
		end = x.End.AsClock()
	}
	return fmt.Sprintf("%s - %s", start, end)
}

func (x TimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
