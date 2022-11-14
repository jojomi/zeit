package zeit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var loc = time.UTC

func TestDate_Begin(t *testing.T) {
	type fields struct {
		value time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "Default",
			fields: fields{
				value: time.Date(2022, time.August, 24, 14, 45, 11, 0, loc),
			},
			want: time.Date(2022, time.August, 24, 0, 0, 0, 0, loc),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Date{
				value: tt.fields.value,
			}
			assert.Equalf(t, tt.want, x.Begin(), "Begin()")
		})
	}
}

func TestDate_End(t *testing.T) {
	type fields struct {
		value time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "Default",
			fields: fields{
				value: time.Date(2022, time.August, 24, 14, 45, 11, 0, loc),
			},
			want: time.Date(2022, time.August, 24, 23, 59, 59, 999999999, loc),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Date{
				value: tt.fields.value,
			}
			assert.Equalf(t, tt.want, x.End(), "End()")
		})
	}
}

func TestDate_Equals(t *testing.T) {
	type fields struct {
		value time.Time
	}
	type args struct {
		o Date
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "match",
			fields: fields{
				value: time.Date(2022, time.August, 24, 14, 45, 11, 0, loc),
			},
			args: args{
				o: Date{
					value: time.Date(2022, time.August, 24, 7, 12, 56, 0, loc),
				},
			},
			want: true,
		},
		{
			name: "no match",
			fields: fields{
				value: time.Date(2022, time.August, 24, 14, 45, 11, 0, loc),
			},
			args: args{
				o: Date{
					value: time.Date(2022, time.September, 24, 7, 12, 56, 0, loc),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Date{
				value: tt.fields.value,
			}
			assert.Equalf(t, tt.want, x.Equals(tt.args.o), "Equals(%v)", tt.args.o)
		})
	}
}

func TestDate_Noon(t *testing.T) {
	type fields struct {
		value time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "Default",
			fields: fields{
				value: time.Date(2022, time.August, 24, 14, 45, 11, 0, loc),
			},
			want: time.Date(2022, time.August, 24, 12, 0, 0, 0, loc),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Date{
				value: tt.fields.value,
			}
			assert.Equalf(t, tt.want, x.Noon(), "Noon()")
		})
	}
}

func TestDate_String(t *testing.T) {
	type fields struct {
		value time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Default",
			fields: fields{
				value: time.Date(2022, time.August, 24, 14, 45, 11, 0, loc),
			},
			want: "2022-08-24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Date{
				value: tt.fields.value,
			}
			assert.Equalf(t, tt.want, x.String(), "String()")
		})
	}
}

func TestNewDateFromTime(t *testing.T) {
	type args struct {
		input time.Time
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		{
			name: "Default",
			args: args{
				input: time.Date(2022, time.August, 24, 14, 45, 11, 0, loc),
			},
			want: Date{
				value: time.Date(2022, time.August, 24, 14, 45, 11, 0, loc),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewDateFromTime(tt.args.input), "NewDateFromTime(%v)", tt.args.input)
		})
	}
}
