package infrastructure

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvToStruct(t *testing.T) {
	const envValue = "dummy value"
	type args struct {
		prefix string
		dest   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "args.dest is set value from environment variable when envconfig tag value is match",
			args: args{
				prefix: "",
				dest: &struct {
					Value string `envconfig:"test_env"`
				}{},
			},
			want: &struct {
				Value string `envconfig:"test_env"`
			}{
				Value: envValue,
			},
			wantErr: assert.NoError,
		},
		{
			name: "args.dest is set value from environment variable when envconfig tag and prefix value is match",
			args: args{
				prefix: "test",
				dest: &struct {
					Value string `envconfig:"env"`
				}{},
			},
			want: &struct {
				Value string `envconfig:"env"`
			}{
				Value: envValue,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Return error when envconfig field value is invalid",
			args: args{
				prefix: "",
				dest: &struct {
					Value string `envconfig:"nothing" required:"true"`
				}{},
			},
			want: &struct {
				Value string `envconfig:"nothing" required:"true"`
			}{},
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "failed to load environment variables: "
				return assert.Contains(tt, e.Error(), want, "LoadEnvToStruct() error = %v, wantErr %v", e, want)
			},
		},
	}
	os.Setenv("TEST_ENV", envValue)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := LoadEnvToStruct(tt.args.prefix, tt.args.dest)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, tt.args.dest, "LoadEnvToStruct() got = %v, want %v", tt.args.dest, tt.want)
		})
	}
}

func TestLoadEnvString(t *testing.T) {
	const keyName = "TEST_VALUE"
	const testValue = "testvalue"
	type args struct {
		key          string
		defaultValue *string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "Return value from environment variable",
			args: args{
				key:          keyName,
				defaultValue: func() *string { value := testValue; return &value }(),
			},
			want: func() *string { value := testValue; return &value }(),
		},
		{
			name: "Return default value when environment variable is empty",
			args: args{
				key:          keyName,
				defaultValue: func() *string { value := testValue; return &value }(),
			},
			want: func() *string { value := testValue; return &value }(),
		},
	}
	os.Setenv(keyName, testValue)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadEnvString(tt.args.key, tt.args.defaultValue)
			assert.Equal(t, tt.want, got, "LoadEnvString() = %v, want %v", got, tt.want)
		})
	}
	os.Setenv(keyName, "")
}
