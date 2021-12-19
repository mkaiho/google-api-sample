package gbpapi

import "errors"

type AccountID int

func ParseAccountID(value int) (*AccountID, error) {
	id := AccountID(value)
	if err := id.validate(); err != nil {
		return nil, err
	}
	return &id, nil
}

func (id AccountID) Int() int {
	return int(id)
}

func (id AccountID) validate() error {
	if id <= 0 {
		return errors.New("invalid account ID")
	}

	return nil
}
