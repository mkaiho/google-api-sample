package gbpapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAccountID(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		args    args
		want    *AccountID
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return account ID",
			args: args{
				value: 1,
			},
			want:    func() *AccountID { value := AccountID(1); return &value }(),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when value is 0",
			args: args{
				value: 0,
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "invalid account ID"
				return assert.EqualError(tt, e, want, "ParseAccountID() error = %v, wantErr %v", e, want)
			},
		},
		{
			name: "Return error when value is less than 0",
			args: args{
				value: -1,
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "invalid account ID"
				return assert.EqualError(tt, e, want, "ParseAccountID() error = %v, wantErr %v", e, want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAccountID(tt.args.value)
			if tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got, "ParseAccountID() = %v, want %v", got, tt.want)
		})
	}
}

func TestAccountID_Int(t *testing.T) {
	tests := []struct {
		name string
		id   AccountID
		want int
	}{
		{
			name: "Return int value",
			id:   AccountID(1),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.id.Int()
			assert.Equal(t, tt.want, got, "AccountID.Int() = %v, want %v", got, tt.want)
		})
	}
}
