package zeit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
func TestIsZeitbereich(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// good
		{
			name: "long variant",
			args: args{
				input: "9:00 -  15:45",
			},
			want: true,
		},
		{
			name: "short variant",
			args: args{
				input: "900-1545",
			},
			want: true,
		},
		// bad
		{
			name: "missing dash",
			args: args{
				input: "9:00 15:45",
			},
			want: false,
		},
		{
			name: "invalid start",
			args: args{
				input: "25:00-15:45",
			},
			want: false,
		},
		{
			name: "invalid end",
			args: args{
				input: "07:25 - 14:70",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsZeitbereich(tt.args.input), "IsZeitbereich(%v)", tt.args.input)
		})
	}
}
*/

func TestNewTimeParsed(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    Time
		wantErr assert.ErrorAssertionFunc
	}{
		// good
		{
			name: "Valid full",
			args: args{
				input: "14:15",
			},
			want: Time{
				Hours:   14,
				Minutes: 15,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Valid short (0 padded)",
			args: args{
				input: "07:30",
			},
			want: Time{
				Hours:   7,
				Minutes: 30,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Valid short (not 0 padded)",
			args: args{
				input: "9:00",
			},
			want: Time{
				Hours:   9,
				Minutes: 0,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Without colon",
			args: args{
				input: "1445",
			},
			want: Time{
				Hours:   14,
				Minutes: 45,
			},
			wantErr: assert.NoError,
		},
		{
			name: "zero hours",
			args: args{
				input: "055",
			},
			want: Time{
				Hours:   0,
				Minutes: 55,
			},
			wantErr: assert.NoError,
		},
		// bad
		{
			name: "hours invalid",
			args: args{
				input: "2450",
			},
			want:    Time{},
			wantErr: assert.Error,
		},
		{
			name: "hours invalid (negative)",
			args: args{
				input: "-5:34",
			},
			want:    Time{},
			wantErr: assert.Error,
		},
		{
			name: "minutes invalid",
			args: args{
				input: "4:60",
			},
			want:    Time{},
			wantErr: assert.Error,
		},
		{
			name: "minutes invalid (negative)",
			args: args{
				input: "11:-30",
			},
			want:    Time{},
			wantErr: assert.Error,
		},
		{
			name: "format invalid",
			args: args{
				input: "17.21",
			},
			want:    Time{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeParsed(tt.args.input)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTimeParsed(%v)", tt.args.input)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTimeParsed(%v)", tt.args.input)
		})
	}
}

func TestNewTime(t *testing.T) {
	type args struct {
		hours   int
		minutes int
	}
	tests := []struct {
		name    string
		args    args
		want    Time
		wantErr assert.ErrorAssertionFunc
	}{
		// good
		{
			name: "making new time",
			args: args{
				hours:   20,
				minutes: 15,
			},
			want: Time{
				Hours:   20,
				Minutes: 15,
			},
			wantErr: assert.NoError,
		},
		// bad
		{
			name: "bad hours",
			args: args{
				hours:   -1,
				minutes: 15,
			},
			want:    Time{},
			wantErr: assert.Error,
		},
		{
			name: "bad hours II",
			args: args{
				hours:   24,
				minutes: 15,
			},
			want:    Time{},
			wantErr: assert.Error,
		},
		{
			name: "bad minutes",
			args: args{
				hours:   12,
				minutes: -1,
			},
			want:    Time{},
			wantErr: assert.Error,
		},
		{
			name: "bad minutes II",
			args: args{
				hours:   12,
				minutes: 60,
			},
			want:    Time{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTime(tt.args.hours, tt.args.minutes)
			if !tt.wantErr(t, err, fmt.Sprintf("NewTime(%v, %v)", tt.args.hours, tt.args.minutes)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewTime(%v, %v)", tt.args.hours, tt.args.minutes)
		})
	}
}
