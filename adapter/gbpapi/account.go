package gbpapi

import (
	"errors"

	"github.com/mkaiho/google-api-sample/util"
)

type AccountID string

func ParseAccountID(value string) (*AccountID, error) {
	id := AccountID(value)
	if err := id.validate(); err != nil {
		return nil, err
	}
	return &id, nil
}

func (id AccountID) String() string {
	return string(id)
}

func (id AccountID) validate() error {
	if util.IsEmptyString(id.String()) {
		return errors.New("invalid account ID")
	}

	return nil
}
