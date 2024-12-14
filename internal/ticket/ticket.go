package ticket

import (
	"errors"
	"regexp"
)

// ErrInvalidFormat is returned when ticket id is invalid
var ErrInvalidFormat = errors.New("invalid ticket format")

// ID represents ticket id
type ID string

// Format represents ticket format
type Format string

const (
	// AlphanumericFormat represents alphanumeric ticket format e.g. FEAT-123
	AlphanumericFormat Format = "^[a-zA-Z]+-[0-9]+$"
	// NumericFormat represents numeric ticket format e.g. 123
	NumericFormat Format = "^[0-9]+$"
)

// String returns string representation of ticket id
func (t ID) String() string {
	return string(t)
}

// ParseID parses ticket id from string
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
