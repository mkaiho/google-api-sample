package infrastructure

import (
	"errors"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/mkaiho/google-api-sample/util"
)

/** Load environment variables and convert to struct **/
func LoadEnvToStruct(prefix string, dest interface{}) error {
	if dest == nil {
		return errors.New("dest was expected non-nil pointer, but not")
	}
	if !util.IsPointer(dest) {
		return errors.New("dest was expected pointer, but not")
	}
	if err := envconfig.Process(prefix, dest); err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	return nil
}

/** Load environment variables and convert to struct **/
func LoadEnvString(key string, defaultValue *string) *string {
	if value := os.Getenv(key); !util.IsEmptyString(value) {
		return &value
	}

	return defaultValue
}
