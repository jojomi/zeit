package zeit

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Duration is an immutable description of a minute-scale time duration.
type Duration struct {
	minutes int
}

// NewDuration creates a new Duration.
func NewDuration() Duration {
	return NewDurationFromMinutes(0)
}

// NewDurationFromTime creates a new Duration from a time.Duration.
func NewDurationFromTime(duration time.Duration) Duration {
	return Duration{
		minutes: int(duration.Minutes()),
	}
}

// NewDurationFromTimes creates a Duration from two Time's. Assumed to be forward in time, a < b.
func NewDurationFromTimes(a, b Time) Duration {
	minutes := b.MinutesSinceMidnight() - a.MinutesSinceMidnight()

	// over midnight?
	if b.Before(&a) {
		minutes = (24*60 - b.MinutesSinceMidnight()) + a.MinutesSinceMidnight()
	}

	return Duration{
		minutes: minutes,
	}
}

// NewDurationFromMinutes creates a Duration from a given number of minutes.
func NewDurationFromMinutes(minutes int) Duration {
	return Duration{
		minutes: minutes,
	}
}

// NewDurationFromHours creates a Duration from a given number of hours.
func NewDurationFromHours(hours float64) Duration {
	return NewDurationFromMinutes(int(hours * 60))
}

// NewDurationFromString creates a Duration from a given string.
func NewDurationFromString(input string) (Duration, error) {
	d := NewDuration()
	err := d.Parse(input)
	return d, err
}

// MustNewDurationFromString creates a Duration from a given string, panicking if it fails.
func MustNewDurationFromString(input string) Duration {
	d, err := NewDurationFromString(input)
	if err != nil {
		panic(err)
	}
	return d
}

// Parse sets this Duration to the interpretation of the given string.
func (x *Duration) Parse(input string) error {
	re := regexp.MustCompile(`^(\d+)(:(\d{2}))?h?$`)
	if !re.MatchString(input) {
		return fmt.Errorf("could not parse duration from %s", input)
	}

	parts := re.FindStringSubmatch(input)

	var (
		hours   = 0
		minutes = 0
		err     error
	)

	hoursString := parts[1]
	if hoursString != "" {
		hours, err = strconv.Atoi(hoursString)
		if err != nil {
			return fmt.Errorf("could not parse duration hours from %s", hoursString)
		}
	}

	minutesString := parts[3]
	if minutesString != "" {
		minutes, err = strconv.Atoi(minutesString)
		if err != nil {
			return fmt.Errorf("could not parse duration minutes from %s", minutesString)
		}
	}

	if minutes < 0 || minutes > 59 {
		return fmt.Errorf("could not parse duration minutes value from %s", minutesString)
	}

	if minutes < 0 || minutes > 59 {
		return fmt.Errorf("could not parse duration minutes value from %s", minutesString)
	}

	x.minutes = hours*60 + minutes
	return nil
}

// Hours returns the number of full hours in this Duration.
func (x *Duration) Hours() int {
	return x.minutes / 60
}

// Minutes returns the number of minutes in this Duration, without full hours.
func (x *Duration) Minutes() int {
	return x.intAbs(x.minutes) % 60
}

// IsPositive returns true if the Duration is positive.
func (x Duration) IsPositive() bool {
	return x.minutes > 0
}

// IsNegative returns true if the Duration is negative.
func (x Duration) IsNegative() bool {
	return x.minutes < 0
}

// AsTime returns this Duration as time.Duration.
func (x Duration) AsTime() time.Duration {
	return time.Duration(int64(x.AsMinutes()) * int64(time.Minute))
}

// AsHours returns this Duration as possibly fractional hours.
func (x Duration) AsHours() float64 {
	return float64(x.minutes) / 60
}

// AsMinutes returns this Duration as total minutes.
func (x Duration) AsMinutes() int {
	return x.minutes
}

// AsUnicode returns the closest unicode clock to this Duration.
func (x Duration) AsUnicode() string {
	charMap := map[string]string{
		"0:00":  "🕛",
		"0:30":  "🕧",
		"1:00":  "🕐",
		"1:30":  "🕜",
		"2:00":  "🕑",
		"2:30":  "🕝",
		"3:00":  "🕒",
		"3:30":  "🕞",
		"4:00":  "🕓",
		"4:30":  "🕟",
		"5:00":  "🕔",
		"5:30":  "🕠",
		"6:00":  "🕕",
		"6:30":  "🕡",
		"7:00":  "🕖",
		"7:30":  "🕢",
		"8:00":  "🕗",
		"8:30":  "🕣",
		"9:00":  "🕘",
		"9:30":  "🕤",
		"10:00": "🕙",
		"10:30": "🕥",
		"11:00": "🕚",
		"11:30": "🕦",
	}

	v := x.RoundedBy(30)
	v.minutes = v.minutes % (12 * 60)

	char, ok := charMap[v.AsRawClock()]
	if !ok {
		return "🕛"
	}
	return char
}

func (x Duration) AsClock() string {
	return x.AsRawClock() + "h"
}

func (x Duration) AsAbsClock() string {
	return x.AsAbsRawClock() + "h"
}

func (x Duration) AsSignedClock() string {
	return x.sign() + x.AsAbsClock()
}

func (x Duration) minimalSign() string {
	if x.IsNegative() {
		return "-"
	}
	return ""
}

func (x Duration) sign() string {
	if x.IsNegative() {
		return "-"
	}
	return "+"
}

func (x Duration) AsRawClock() string {
	return x.Format("3:04")
}

func (x Duration) AsAbsRawClock() string {
	return x.Format("|3:04")
}

func (x Duration) AsFrac(lang *language.Tag) string {
	return x.AsRawFrac(lang) + "h"
}

func (x Duration) AsRawFrac(lang *language.Tag) string {
	var (
		result       string
		formatString = "%.2f"
	)
	if lang != nil {
		loc := message.NewPrinter(*lang)
		result = loc.Sprintf(formatString, x.AsHours())
	} else {
		result = fmt.Sprintf(formatString, x.AsHours())
	}

	// cut trailing zeros AND trailing separator if it remains
	result = x.fixHourString(result)
	return result
}

func (x Duration) Add(other Duration) Duration {
	return NewDurationFromMinutes(x.AsMinutes() + other.AsMinutes())
}

func (x Duration) Sub(other Duration) Duration {
	return NewDurationFromMinutes(x.AsMinutes() - other.AsMinutes())
}

func (x Duration) Equals(other Duration) bool {
	return x.AsMinutes() == other.AsMinutes()
}

func (x Duration) IsZero() bool {
	return x.minutes == 0
}

func (x Duration) String() string {
	return x.AsClock()
}

func (x Duration) fixHourString(input string) string {
	r := regexp.MustCompile(`[^0-9]?0{1,2}$`)
	input = r.ReplaceAllString(input, "")
	return input
}

func (x Duration) longFrac(input string, length int) string {
	r := regexp.MustCompile(`[^0-9]\d+$`)
	input = r.ReplaceAllStringFunc(input, func(input string) string {
		for {
			if len(input) >= length {
				return input
			}
			input = input + "0"
		}
	})
	return input
}

func (x *Duration) intAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// RoundedBy returns a new Duration rounded to the next value given by a number of minutes or a fraction of on hour (for values <= 1).
// Make sure to only use values that cleanly split the 60 minutes of an hour like 15 (=0.25), 20 (=0.33), 30 (=0.5) or 60 (= 1).
func (x Duration) RoundedBy(rounding float64) Duration {
	if rounding <= 1 {
		rounding *= 60
	}

	minutes := x.AsMinutes()
	minutes = int(math.Round(float64(minutes)/rounding) * rounding)

	return NewDurationFromMinutes(minutes)
}

// RoundedDownBy returns a new Duration rounded down to the next value given by a number of minutes or a fraction of on hour (for values <= 1).
// Make sure to only use values that cleanly split the 60 minutes of an hour like 15 (=0.25), 20 (=0.33), 30 (=0.5) or 60 (= 1).
// Also see RoundedBy.
func (x Duration) RoundedDownBy(rounding float64) Duration {
	if rounding <= 1 {
		rounding *= 60
	}

	minutes := x.AsMinutes()
	minutes = int(math.Floor(float64(minutes)/rounding) * rounding)

	return NewDurationFromMinutes(minutes)
}

// RoundedUpBy returns a new Duration rounded down to the next value given by a number of minutes or a fraction of on hour (for values <= 1).
// Also see RoundedBy.
func (x Duration) RoundedUpBy(rounding float64) Duration {
	if rounding <= 1 {
		rounding *= 60
	}

	minutes := x.AsMinutes()
	minutes = int(math.Ceil(float64(minutes)/rounding) * rounding)

	return NewDurationFromMinutes(minutes)
}

// Format allows to format a duration similar to time.Format.
// Valid elements are
//
//	3: hours part
//	|3: absolute hours part
//	+3: hours part with forced sign indicator (also show +)
//	04: minutes part with 0 padding to 2 characters if necessary
//	4: minutes part
//	07: minutes part as percentage with maximum 2 characters
//	0-: minutes part as percentage with maximum 2 characters, but minimal length too
//	184: total minutes
func (x Duration) Format(layout string) string {
	m := x.Minutes()
	h := x.Hours()
	hAbs := x.intAbs(h)
	frac := strings.Replace(fmt.Sprintf("%.2f", float64(m)/60.0), "0.", "", 1)

	rep := strings.NewReplacer(
		"+3", fmt.Sprintf("%s%d", x.sign(), hAbs),
		"|3", fmt.Sprintf("%d", hAbs),
		"3", fmt.Sprintf("%s%d", x.minimalSign(), hAbs),
		"04", fmt.Sprintf("%02d", m),
		"4", strconv.Itoa(m),
		"0-", x.fixHourString(frac),
		"07", x.longFrac(frac, 2),
		"184", strconv.Itoa(x.minutes),
	)
	return rep.Replace(layout)
}

func (x Duration) shortenFrac(input string) string {
	output := input
	for {
		if !strings.HasSuffix(output, "0") {
			break
		}
		output = output[0 : len(output)-1]
	}
	if strings.HasSuffix(output, ".") {
		output = output[0 : len(output)-1]
	}
	return output
}

func (x Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

func (x Duration) UnmarshalJSON(data []byte) error {
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return err
	}

	return x.Parse(strVal)
}
