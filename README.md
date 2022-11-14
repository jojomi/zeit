# zeit

`zeit` does contain a few time abstractions that might be useful in other projects.

[![Go Reference](https://pkg.go.dev/badge/github.com/jojomi/zeit.svg)](https://pkg.go.dev/github.com/jojomi/zeit)

## Key structs

* **[Date](https://pkg.go.dev/github.com/jojomi/zeit#Date)**: Date is a single day with no relevant time of day.
 (full source: [date.go](date.go))
* **[Time](https://pkg.go.dev/github.com/jojomi/zeit#Time)**: Time denotes a time within a day.
 (full source: [time.go](time.go))
* **[Duration](https://pkg.go.dev/github.com/jojomi/zeit#Duration)**: Duration is an immutable description of a minute-scale time duration.
 (full source: [duration.go](duration.go))
* **[Month](https://pkg.go.dev/github.com/jojomi/zeit#Month)**: Month is a month in a year. Supports navigating back and forth as well as printing in different styles.
 (full source: [month.go](month.go))
* **[TimeRange](https://pkg.go.dev/github.com/jojomi/zeit#TimeRange)**: TimeRange is a pair of start and end time while both are optional.
 (full source: [time_range.go](time_range.go))


## Tests

``` shell
go test -count=1 -v ./...
ok  	github.com/jojomi/zeit	0.350s

```

<details>
  <summary>All test results</summary>

  ```
  === RUN   Test_parseWithDayDiff
--- PASS: Test_parseWithDayDiff (0.00s)
=== RUN   TestDate_Begin
=== RUN   TestDate_Begin/Default
--- PASS: TestDate_Begin (0.00s)
    --- PASS: TestDate_Begin/Default (0.00s)
=== RUN   TestDate_End
=== RUN   TestDate_End/Default
--- PASS: TestDate_End (0.00s)
    --- PASS: TestDate_End/Default (0.00s)
=== RUN   TestDate_Equals
=== RUN   TestDate_Equals/match
=== RUN   TestDate_Equals/no_match
--- PASS: TestDate_Equals (0.00s)
    --- PASS: TestDate_Equals/match (0.00s)
    --- PASS: TestDate_Equals/no_match (0.00s)
=== RUN   TestDate_Noon
=== RUN   TestDate_Noon/Default
--- PASS: TestDate_Noon (0.00s)
    --- PASS: TestDate_Noon/Default (0.00s)
=== RUN   TestDate_String
=== RUN   TestDate_String/Default
--- PASS: TestDate_String (0.00s)
    --- PASS: TestDate_String/Default (0.00s)
=== RUN   TestNewDateFromTime
=== RUN   TestNewDateFromTime/Default
--- PASS: TestNewDateFromTime (0.00s)
    --- PASS: TestNewDateFromTime/Default (0.00s)
=== RUN   TestDuration_AsClock
=== RUN   TestDuration_AsClock/basic
=== RUN   TestDuration_AsClock/negative
--- PASS: TestDuration_AsClock (0.00s)
    --- PASS: TestDuration_AsClock/basic (0.00s)
    --- PASS: TestDuration_AsClock/negative (0.00s)
=== RUN   TestDuration_AsHours
=== RUN   TestDuration_AsHours/basic
=== RUN   TestDuration_AsHours/negative
--- PASS: TestDuration_AsHours (0.00s)
    --- PASS: TestDuration_AsHours/basic (0.00s)
    --- PASS: TestDuration_AsHours/negative (0.00s)
=== RUN   TestDuration_AsMinutes
=== RUN   TestDuration_AsMinutes/basic
=== RUN   TestDuration_AsMinutes/negative
--- PASS: TestDuration_AsMinutes (0.00s)
    --- PASS: TestDuration_AsMinutes/basic (0.00s)
    --- PASS: TestDuration_AsMinutes/negative (0.00s)
=== RUN   TestDuration_AsRawClock
=== RUN   TestDuration_AsRawClock/basic
--- PASS: TestDuration_AsRawClock (0.00s)
    --- PASS: TestDuration_AsRawClock/basic (0.00s)
=== RUN   TestNewDurationFromHours
=== RUN   TestNewDurationFromHours/basic
=== RUN   TestNewDurationFromHours/negative
--- PASS: TestNewDurationFromHours (0.00s)
    --- PASS: TestNewDurationFromHours/basic (0.00s)
    --- PASS: TestNewDurationFromHours/negative (0.00s)
=== RUN   TestNewDurationFromMinutes
=== RUN   TestNewDurationFromMinutes/basic
=== RUN   TestNewDurationFromMinutes/negative
--- PASS: TestNewDurationFromMinutes (0.00s)
    --- PASS: TestNewDurationFromMinutes/basic (0.00s)
    --- PASS: TestNewDurationFromMinutes/negative (0.00s)
=== RUN   TestDuration_AsFrac
=== RUN   TestDuration_AsFrac/basic
=== RUN   TestDuration_AsFrac/basic_no_language
--- PASS: TestDuration_AsFrac (0.00s)
    --- PASS: TestDuration_AsFrac/basic (0.00s)
    --- PASS: TestDuration_AsFrac/basic_no_language (0.00s)
=== RUN   TestDuration_AsRawFrac
=== RUN   TestDuration_AsRawFrac/basic
--- PASS: TestDuration_AsRawFrac (0.00s)
    --- PASS: TestDuration_AsRawFrac/basic (0.00s)
=== RUN   TestDuration_IsZero
=== RUN   TestDuration_IsZero/positive_non-zero
=== RUN   TestDuration_IsZero/negative_non-zero
=== RUN   TestDuration_IsZero/zero
--- PASS: TestDuration_IsZero (0.00s)
    --- PASS: TestDuration_IsZero/positive_non-zero (0.00s)
    --- PASS: TestDuration_IsZero/negative_non-zero (0.00s)
    --- PASS: TestDuration_IsZero/zero (0.00s)
=== RUN   TestDuration_Minutes
=== RUN   TestDuration_Minutes/basic
=== RUN   TestDuration_Minutes/negative
--- PASS: TestDuration_Minutes (0.00s)
    --- PASS: TestDuration_Minutes/basic (0.00s)
    --- PASS: TestDuration_Minutes/negative (0.00s)
=== RUN   TestDuration_AsSignedClock
=== RUN   TestDuration_AsSignedClock/negative
=== RUN   TestDuration_AsSignedClock/positive
--- PASS: TestDuration_AsSignedClock (0.00s)
    --- PASS: TestDuration_AsSignedClock/negative (0.00s)
    --- PASS: TestDuration_AsSignedClock/positive (0.00s)
=== RUN   TestDuration_RoundedBy
=== RUN   TestDuration_RoundedBy/round_to_full_hour
=== RUN   TestDuration_RoundedBy/round_down_to_full_hour
=== RUN   TestDuration_RoundedBy/round_to_half_hour
=== RUN   TestDuration_RoundedBy/round_to_half_hour#01
=== RUN   TestDuration_RoundedBy/round_down_big_number_of_minutes
=== RUN   TestDuration_RoundedBy/round_down_with_fraction
--- PASS: TestDuration_RoundedBy (0.00s)
    --- PASS: TestDuration_RoundedBy/round_to_full_hour (0.00s)
    --- PASS: TestDuration_RoundedBy/round_down_to_full_hour (0.00s)
    --- PASS: TestDuration_RoundedBy/round_to_half_hour (0.00s)
    --- PASS: TestDuration_RoundedBy/round_to_half_hour#01 (0.00s)
    --- PASS: TestDuration_RoundedBy/round_down_big_number_of_minutes (0.00s)
    --- PASS: TestDuration_RoundedBy/round_down_with_fraction (0.00s)
=== RUN   TestDuration_RoundedDownBy
=== RUN   TestDuration_RoundedDownBy/round_down
=== RUN   TestDuration_RoundedDownBy/round_down_to_zero
=== RUN   TestDuration_RoundedDownBy/round_down_by_factor
--- PASS: TestDuration_RoundedDownBy (0.00s)
    --- PASS: TestDuration_RoundedDownBy/round_down (0.00s)
    --- PASS: TestDuration_RoundedDownBy/round_down_to_zero (0.00s)
    --- PASS: TestDuration_RoundedDownBy/round_down_by_factor (0.00s)
=== RUN   TestDuration_RoundedUpBy
=== RUN   TestDuration_RoundedUpBy/round_up
=== RUN   TestDuration_RoundedUpBy/round_up_from_zero
=== RUN   TestDuration_RoundedUpBy/round_up_by_factor
--- PASS: TestDuration_RoundedUpBy (0.00s)
    --- PASS: TestDuration_RoundedUpBy/round_up (0.00s)
    --- PASS: TestDuration_RoundedUpBy/round_up_from_zero (0.00s)
    --- PASS: TestDuration_RoundedUpBy/round_up_by_factor (0.00s)
=== RUN   TestDuration_IsNegative
=== RUN   TestDuration_IsNegative/negative_value
=== RUN   TestDuration_IsNegative/poitive_value
=== RUN   TestDuration_IsNegative/zero_value
--- PASS: TestDuration_IsNegative (0.00s)
    --- PASS: TestDuration_IsNegative/negative_value (0.00s)
    --- PASS: TestDuration_IsNegative/poitive_value (0.00s)
    --- PASS: TestDuration_IsNegative/zero_value (0.00s)
=== RUN   TestDuration_IsZero1
=== RUN   TestDuration_IsZero1/zero
=== RUN   TestDuration_IsZero1/non-zero
--- PASS: TestDuration_IsZero1 (0.00s)
    --- PASS: TestDuration_IsZero1/zero (0.00s)
    --- PASS: TestDuration_IsZero1/non-zero (0.00s)
=== RUN   TestDuration_IsPositive
=== RUN   TestDuration_IsPositive/positive_value
=== RUN   TestDuration_IsPositive/negative_value
=== RUN   TestDuration_IsPositive/zero_value
--- PASS: TestDuration_IsPositive (0.00s)
    --- PASS: TestDuration_IsPositive/positive_value (0.00s)
    --- PASS: TestDuration_IsPositive/negative_value (0.00s)
    --- PASS: TestDuration_IsPositive/zero_value (0.00s)
=== RUN   TestDuration_String
=== RUN   TestDuration_String/basic_test
--- PASS: TestDuration_String (0.00s)
    --- PASS: TestDuration_String/basic_test (0.00s)
=== RUN   TestDuration_Equals
=== RUN   TestDuration_Equals/same_value
=== RUN   TestDuration_Equals/different_value
--- PASS: TestDuration_Equals (0.00s)
    --- PASS: TestDuration_Equals/same_value (0.00s)
    --- PASS: TestDuration_Equals/different_value (0.00s)
=== RUN   TestDuration_Add
=== RUN   TestDuration_Add/add
--- PASS: TestDuration_Add (0.00s)
    --- PASS: TestDuration_Add/add (0.00s)
=== RUN   TestDuration_Sub
=== RUN   TestDuration_Sub/sub
--- PASS: TestDuration_Sub (0.00s)
    --- PASS: TestDuration_Sub/sub (0.00s)
=== RUN   TestNewDuration
=== RUN   TestNewDuration/simple
--- PASS: TestNewDuration (0.00s)
    --- PASS: TestNewDuration/simple (0.00s)
=== RUN   TestDuration_Format
=== RUN   TestDuration_Format/raw_clock_formatted
=== RUN   TestDuration_Format/raw_frac_formatted
=== RUN   TestDuration_Format/raw_frac_shortened
=== RUN   TestDuration_Format/abs_formatted
=== RUN   TestDuration_Format/pre_and_post_text
=== RUN   TestDuration_Format/sign
=== RUN   TestDuration_Format/forced_sign
--- PASS: TestDuration_Format (0.00s)
    --- PASS: TestDuration_Format/raw_clock_formatted (0.00s)
    --- PASS: TestDuration_Format/raw_frac_formatted (0.00s)
    --- PASS: TestDuration_Format/raw_frac_shortened (0.00s)
    --- PASS: TestDuration_Format/abs_formatted (0.00s)
    --- PASS: TestDuration_Format/pre_and_post_text (0.00s)
    --- PASS: TestDuration_Format/sign (0.00s)
    --- PASS: TestDuration_Format/forced_sign (0.00s)
=== RUN   TestDuration_Parse
=== RUN   TestDuration_Parse/basic
=== RUN   TestDuration_Parse/hours_only
=== RUN   TestDuration_Parse/without_suffix
=== RUN   TestDuration_Parse/invalid
=== RUN   TestDuration_Parse/invalid_minutes
=== RUN   TestDuration_Parse/invalid_negative_minutes
=== RUN   TestDuration_Parse/invalid_single_digit_minutes
=== RUN   TestDuration_Parse/big_hours_duration
--- PASS: TestDuration_Parse (0.00s)
    --- PASS: TestDuration_Parse/basic (0.00s)
    --- PASS: TestDuration_Parse/hours_only (0.00s)
    --- PASS: TestDuration_Parse/without_suffix (0.00s)
    --- PASS: TestDuration_Parse/invalid (0.00s)
    --- PASS: TestDuration_Parse/invalid_minutes (0.00s)
    --- PASS: TestDuration_Parse/invalid_negative_minutes (0.00s)
    --- PASS: TestDuration_Parse/invalid_single_digit_minutes (0.00s)
    --- PASS: TestDuration_Parse/big_hours_duration (0.00s)
=== RUN   TestNewTimeParsed
=== RUN   TestNewTimeParsed/Valid_full
=== RUN   TestNewTimeParsed/Valid_short_(0_padded)
=== RUN   TestNewTimeParsed/Valid_short_(not_0_padded)
=== RUN   TestNewTimeParsed/Without_colon
=== RUN   TestNewTimeParsed/zero_hours
=== RUN   TestNewTimeParsed/hours_invalid
=== RUN   TestNewTimeParsed/hours_invalid_(negative)
=== RUN   TestNewTimeParsed/minutes_invalid
=== RUN   TestNewTimeParsed/minutes_invalid_(negative)
=== RUN   TestNewTimeParsed/format_invalid
--- PASS: TestNewTimeParsed (0.00s)
    --- PASS: TestNewTimeParsed/Valid_full (0.00s)
    --- PASS: TestNewTimeParsed/Valid_short_(0_padded) (0.00s)
    --- PASS: TestNewTimeParsed/Valid_short_(not_0_padded) (0.00s)
    --- PASS: TestNewTimeParsed/Without_colon (0.00s)
    --- PASS: TestNewTimeParsed/zero_hours (0.00s)
    --- PASS: TestNewTimeParsed/hours_invalid (0.00s)
    --- PASS: TestNewTimeParsed/hours_invalid_(negative) (0.00s)
    --- PASS: TestNewTimeParsed/minutes_invalid (0.00s)
    --- PASS: TestNewTimeParsed/minutes_invalid_(negative) (0.00s)
    --- PASS: TestNewTimeParsed/format_invalid (0.00s)
=== RUN   TestNewTime
=== RUN   TestNewTime/making_new_time
=== RUN   TestNewTime/bad_hours
=== RUN   TestNewTime/bad_hours_II
=== RUN   TestNewTime/bad_minutes
=== RUN   TestNewTime/bad_minutes_II
--- PASS: TestNewTime (0.00s)
    --- PASS: TestNewTime/making_new_time (0.00s)
    --- PASS: TestNewTime/bad_hours (0.00s)
    --- PASS: TestNewTime/bad_hours_II (0.00s)
    --- PASS: TestNewTime/bad_minutes (0.00s)
    --- PASS: TestNewTime/bad_minutes_II (0.00s)
PASS
ok  	github.com/jojomi/zeit	0.180s
  ```
</details>

### Test Coverage

```
go test -cover -count=1 ./...
ok  	github.com/jojomi/zeit	0.267s	coverage: 44.2% of statements
```

## Latest changes

See the (https://github.com/jojomi/zeit/commits/master)[commits on master].

## Why the name?

"Zeit" is German for "time".
