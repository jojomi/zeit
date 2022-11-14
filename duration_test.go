package zeit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestDuration_AsClock(t *testing.T) {
	type fields struct {
		hours   int
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				minutes: 4*60 + 24,
			},
			want: "4:24h",
		},
		{
			name: "negative",
			fields: fields{
				minutes: -2 * 60,
			},
			want: "-2:00h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.AsClock(), "AsClock()")
		})
	}
}

func TestDuration_AsHours(t *testing.T) {
	type fields struct {
		hours   int
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "basic",
			fields: fields{
				minutes: 15,
			},
			want: 0.25,
		},
		{
			name: "negative",
			fields: fields{
				minutes: -1 * 60,
			},
			want: -1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.AsHours(), "AsHours()")
		})
	}
}

func TestDuration_AsMinutes(t *testing.T) {
	type fields struct {
		hours   int
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "basic",
			fields: fields{
				minutes: 10*60 + 6,
			},
			want: 606,
		},
		{
			name: "negative",
			fields: fields{
				minutes: -2 * 60,
			},
			want: -120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.AsMinutes(), "AsMinutes()")
		})
	}
}

func TestDuration_AsRawClock(t *testing.T) {
	type fields struct {
		hours   int
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				minutes: 25*60 + 15,
			},
			want: "25:15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.AsRawClock(), "AsRawClock()")
		})
	}
}

func TestNewDurationFromHours(t *testing.T) {
	type args struct {
		hours float64
	}
	tests := []struct {
		name string
		args args
		want Duration
	}{
		{
			name: "basic",
			args: args{
				hours: 0.5,
			},
			want: Duration{
				minutes: 30,
			},
		},
		{
			name: "negative",
			args: args{
				hours: -2.75,
			},
			want: Duration{
				minutes: -(2*60 + 45),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewDurationFromHours(tt.args.hours), "NewDurationFromHours(%v)", tt.args.hours)
		})
	}
}

func TestNewDurationFromMinutes(t *testing.T) {
	type args struct {
		minutes int
	}
	tests := []struct {
		name string
		args args
		want Duration
	}{
		{
			name: "basic",
			args: args{
				minutes: 123,
			},
			want: Duration{
				minutes: 2*60 + 3,
			},
		},
		{
			name: "negative",
			args: args{
				minutes: -119,
			},
			want: Duration{
				minutes: -(1*60 + 59),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewDurationFromMinutes(tt.args.minutes), "NewDurationFromMinutes(%v)", tt.args.minutes)
		})
	}
}

func TestDuration_AsFrac(t *testing.T) {
	type fields struct {
		hours   int
		minutes int
	}
	type args struct {
		lang *language.Tag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				minutes: 3*60 + 30,
			},
			args: args{
				lang: &language.German,
			},
			want: "3,5h",
		},
		{
			name: "basic no language",
			fields: fields{
				minutes: 3*60 + 30,
			},
			args: args{
				lang: nil,
			},
			want: "3.5h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.AsFrac(tt.args.lang), "AsFrac(%v)", tt.args.lang)
		})
	}
}

func TestDuration_AsRawFrac(t *testing.T) {
	type fields struct {
		hours   int
		minutes int
	}
	type args struct {
		lang *language.Tag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				minutes: 9*60 + 9,
			},
			args: args{
				lang: &language.German,
			},
			want: "9,15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.AsRawFrac(tt.args.lang), "AsRawFrac(%v)", tt.args.lang)
		})
	}
}

func TestDuration_IsZero(t *testing.T) {
	type fields struct {
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "positive non-zero",
			fields: fields{
				minutes: 10,
			},
			want: false,
		},
		{
			name: "negative non-zero",
			fields: fields{
				minutes: -1,
			},
			want: false,
		},
		{
			name: "zero",
			fields: fields{
				minutes: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.IsZero(), "IsZero()")
		})
	}
}

func TestDuration_Minutes(t *testing.T) {
	type fields struct {
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "basic",
			fields: fields{
				minutes: 150,
			},
			want: 30,
		},
		{
			name: "negative",
			fields: fields{
				minutes: -75,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.Minutes(), "Minutes()")
		})
	}
}

func TestDuration_AsSignedClock(t *testing.T) {
	type fields struct {
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "negative",
			fields: fields{
				minutes: -70,
			},
			want: "-1:10h",
		},
		{
			name: "positive",
			fields: fields{
				minutes: 30,
			},
			want: "+0:30h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.AsSignedClock(), "AsSignedClock()")
		})
	}
}

func TestDuration_RoundedBy(t *testing.T) {
	type fields struct {
		minutes int
	}
	type args struct {
		rounding float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Duration
	}{
		{
			name: "round to full hour",
			fields: fields{
				minutes: 55,
			},
			args: args{
				rounding: 60,
			},
			want: Duration{
				minutes: 60,
			},
		},
		{
			name: "round down to full hour",
			fields: fields{
				minutes: 75,
			},
			args: args{
				rounding: 60,
			},
			want: Duration{
				minutes: 60,
			},
		},
		{
			name: "round to half hour",
			fields: fields{
				minutes: 77,
			},
			args: args{
				rounding: 30,
			},
			want: Duration{
				minutes: 90,
			},
		},
		{
			name: "round to half hour",
			fields: fields{
				minutes: 10,
			},
			args: args{
				rounding: 30,
			},
			want: Duration{
				minutes: 0,
			},
		},
		{
			name: "round down big number of minutes",
			fields: fields{
				minutes: 182,
			},
			args: args{
				rounding: 90,
			},
			want: Duration{
				minutes: 180,
			},
		},

		{
			name: "round down with fraction",
			fields: fields{
				minutes: 78,
			},
			args: args{
				rounding: 0.25,
			},
			want: Duration{
				minutes: 75,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.RoundedBy(tt.args.rounding), "RoundedBy(%v)", tt.args.rounding)
		})
	}
}

func TestDuration_RoundedDownBy(t *testing.T) {
	type fields struct {
		minutes int
	}
	type args struct {
		rounding float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Duration
	}{
		{
			name: "round down",
			fields: fields{
				minutes: 55,
			},
			args: args{
				rounding: 30,
			},
			want: Duration{
				minutes: 30,
			},
		},
		{
			name: "round down to zero",
			fields: fields{
				minutes: 55,
			},
			args: args{
				rounding: 60,
			},
			want: Duration{
				minutes: 0,
			},
		},
		{
			name: "round down by factor",
			fields: fields{
				minutes: 119,
			},
			args: args{
				rounding: 0.25,
			},
			want: Duration{
				minutes: 105,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.RoundedDownBy(tt.args.rounding), "RoundedDownBy(%v)", tt.args.rounding)
		})
	}
}

func TestDuration_RoundedUpBy(t *testing.T) {
	type fields struct {
		minutes int
	}
	type args struct {
		rounding float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Duration
	}{
		{
			name: "round up",
			fields: fields{
				minutes: 68,
			},
			args: args{
				rounding: 20,
			},
			want: Duration{
				minutes: 80,
			},
		},
		{
			name: "round up from zero",
			fields: fields{
				minutes: 0,
			},
			args: args{
				rounding: 20,
			},
			want: Duration{
				minutes: 0,
			},
		},
		{
			name: "round up by factor",
			fields: fields{
				minutes: 1,
			},
			args: args{
				rounding: 0.5,
			},
			want: Duration{
				minutes: 30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.RoundedUpBy(tt.args.rounding), "RoundedUpBy(%v)", tt.args.rounding)
		})
	}
}

func TestDuration_IsNegative(t *testing.T) {
	type fields struct {
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "negative value",
			fields: fields{
				minutes: -14,
			},
			want: true,
		},
		{
			name: "poitive value",
			fields: fields{
				minutes: 1,
			},
			want: false,
		},
		{
			name: "zero value",
			fields: fields{
				minutes: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.IsNegative(), "IsNegative()")
		})
	}
}

func TestDuration_IsZero1(t *testing.T) {
	type fields struct {
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero",
			fields: fields{
				minutes: 0,
			},
			want: true,
		},
		{
			name: "non-zero",
			fields: fields{
				minutes: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.IsZero(), "IsZero()")
		})
	}
}

func TestDuration_IsPositive(t *testing.T) {
	type fields struct {
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "positive value",
			fields: fields{
				minutes: 99923,
			},
			want: true,
		},
		{
			name: "negative value",
			fields: fields{
				minutes: -1,
			},
			want: false,
		},
		{
			name: "zero value",
			fields: fields{
				minutes: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.IsPositive(), "IsPositive()")
		})
	}
}

func TestDuration_String(t *testing.T) {
	type fields struct {
		minutes int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "basic test",
			fields: fields{
				minutes: 64,
			},
			want: "1:04h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.String(), "String()")
		})
	}
}

func TestDuration_Equals(t *testing.T) {
	type fields struct {
		minutes int
	}
	type args struct {
		other Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "same value",
			fields: fields{
				minutes: 3,
			},
			args: args{
				other: NewDurationFromMinutes(3),
			},
			want: true,
		},
		{
			name: "different value",
			fields: fields{
				minutes: 6,
			},
			args: args{
				other: NewDurationFromMinutes(4),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.Equals(tt.args.other), "Equals(%v)", tt.args.other)
		})
	}
}

func TestDuration_Add(t *testing.T) {
	type fields struct {
		minutes int
	}
	type args struct {
		other Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Duration
	}{
		{
			name: "add",
			fields: fields{
				minutes: -10,
			},
			args: args{
				other: NewDurationFromMinutes(24),
			},
			want: Duration{
				minutes: 14,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.Add(tt.args.other), "Add(%v)", tt.args.other)
		})
	}
}

func TestDuration_Sub(t *testing.T) {
	type fields struct {
		minutes int
	}
	type args struct {
		other Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Duration
	}{
		{
			name: "sub",
			fields: fields{
				minutes: 55,
			},
			args: args{
				other: NewDurationFromMinutes(1000),
			},
			want: Duration{
				minutes: -945,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.Sub(tt.args.other), "Sub(%v)", tt.args.other)
		})
	}
}

func TestNewDuration(t *testing.T) {
	tests := []struct {
		name string
		want Duration
	}{
		{
			name: "simple",
			want: Duration{
				minutes: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewDuration(), "NewDuration()")
		})
	}
}

func TestDuration_Format(t *testing.T) {
	type fields struct {
		minutes int
	}
	type args struct {
		layout string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "raw clock formatted",
			fields: fields{
				minutes: 444,
			},
			args: args{
				layout: "3:04",
			},
			want: "7:24",
		},
		{
			name: "raw frac formatted",
			fields: fields{
				minutes: -150,
			},
			args: args{
				layout: "3.07",
			},
			want: "-2.50",
		},
		{
			name: "raw frac shortened",
			fields: fields{
				minutes: -150,
			},
			args: args{
				layout: "3.0-",
			},
			want: "-2.5",
		},
		{
			name: "abs formatted",
			fields: fields{
				minutes: -30,
			},
			args: args{
				layout: "|3,07",
			},
			want: "0,50",
		},
		{
			name: "pre and post text",
			fields: fields{
				minutes: 71,
			},
			args: args{
				layout: "Duration: 3:04 hours",
			},
			want: "Duration: 1:11 hours",
		},
		{
			name: "sign",
			fields: fields{
				minutes: -122,
			},
			args: args{
				layout: "3 hours and 4 minutes",
			},
			want: "-2 hours and 2 minutes",
		},
		{
			name: "forced sign",
			fields: fields{
				minutes: 69,
			},
			args: args{
				layout: "+3:4",
			},
			want: "+1:9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Duration{
				minutes: tt.fields.minutes,
			}
			assert.Equalf(t, tt.want, x.Format(tt.args.layout), "Format(%v)", tt.args.layout)
		})
	}
}

func TestDuration_Parse(t *testing.T) {
	type fields struct {
		minutes int
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "basic",
			fields: fields{
				minutes: 0,
			},
			args: args{
				input: "5:52h",
			},
			want:    5*60 + 52,
			wantErr: assert.NoError,
		},
		{
			name: "hours only",
			fields: fields{
				minutes: 0,
			},
			args: args{
				input: "5h",
			},
			want:    5 * 60,
			wantErr: assert.NoError,
		},
		{
			name: "without suffix",
			fields: fields{
				minutes: 0,
			},
			args: args{
				input: "0:22",
			},
			want:    22,
			wantErr: assert.NoError,
		},
		{
			name:   "invalid",
			fields: fields{},
			args: args{
				input: "3:0:44h",
			},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name: "invalid minutes",
			fields: fields{
				minutes: 30,
			},
			args: args{
				input: "1:77h",
			},
			want:    30,
			wantErr: assert.Error,
		},
		{
			name: "invalid negative minutes",
			fields: fields{
				minutes: 30,
			},
			args: args{
				input: "1:-77h",
			},
			want:    30,
			wantErr: assert.Error,
		},
		{
			name: "invalid single digit minutes",
			fields: fields{
				minutes: 30,
			},
			args: args{
				input: "1:8h",
			},
			want:    30,
			wantErr: assert.Error,
		},
		{
			name: "big hours duration",
			fields: fields{
				minutes: 30,
			},
			args: args{
				input: "1000:00h",
			},
			want:    1000 * 60,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Duration{
				minutes: tt.fields.minutes,
			}
			tt.wantErr(t, x.Parse(tt.args.input), fmt.Sprintf("Parse(%v)", tt.args.input))
			assert.Equalf(t, tt.want, x.minutes, "expected minutes (%v)", tt.want)
		})
	}
}
