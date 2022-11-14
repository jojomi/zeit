package zeit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
func TestParseDate(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    Date
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "today",
			args: args{
				input: "today",
			},
			want:    NewDateFromTime(time.Now().In(loc)),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDate(tt.args.input)
			fmt.Println("got", got, err)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseDate(%v)", tt.args.input)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParseDate(%v)", tt.args.input)
		})
	}
}
*/

func Test_parseWithDayDiff(t *testing.T) {
	type args struct {
		diffDays int
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseWithDayDiff(tt.args.diffDays), "parseWithDayDiff(%v)", tt.args.diffDays)
		})
	}
}
