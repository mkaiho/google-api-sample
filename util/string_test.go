package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyString(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Return true when value is empty",
			args: args{
				v: "",
			},
			want: true,
		},
		{
			name: "Return false when value is not empty",
			args: args{
				v: "dummy",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEmptyString(tt.args.v)
			assert.Equal(t, tt.want, got, "IsEmptyString() = %v, want %v", got, tt.want)
		})
	}
}
