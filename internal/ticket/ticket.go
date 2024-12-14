package ticket

import (
	"errors"
	"regexp"
)

var ErrInvalidFormat = errors.New("invalid ticket format")

type ID string

type Format string

const (
	AlphanumericFormat Format = "^[a-zA-Z]+-[0-9]+$"
	NumericFormat      Format = "^[0-9]+$"
)

func (t ID) String() string {
	return string(t)
}

func ParseID(s string, format Format) (ID, error) {
	re, err := regexp.Compile(string(format))
	if err != nil {
		return "", err
	}
	if !re.MatchString(s) {
		return "", ErrInvalidFormat
	}
	return ID(s), nil
}
