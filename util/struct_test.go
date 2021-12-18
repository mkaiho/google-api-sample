package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPointer(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Return true when value is pointer",
			args: args{
				value: new(interface{}),
			},
			want: true,
		},
		{
			name: "Return false when value is not pointer",
			args: args{
				value: *new(interface{}),
			},
			want: false,
		},
		{
			name: "Return false when value is nil",
			args: args{
				value: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPointer(tt.args.value)
			assert.Equal(t, tt.want, got, "IsPointer() = %v, want %v", got, tt.want)
		})
	}
}
