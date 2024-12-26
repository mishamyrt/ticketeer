package ticket

import (
	"errors"
	"regexp"
)

// ErrInvalidFormat is returned when ticket id is invalid
var ErrInvalidFormat = errors.New("invalid ticket format")

// ID represents ticket id
type ID string

// IDFormat represents ticket format
type IDFormat string

// alphanumeric, alphanumeric-small, alphanumeric-caps, numeric
const (
	// AlphanumericFormat represents alphanumeric ticket format e.g. FEAT-123
	AlphanumericFormat IDFormat = "^([a-zA-Z]+-[0-9]+)$"

	// AlphanumericSmallFormat represents alphanumeric small ticket format e.g. feat-123
	AlphanumericSmallFormat IDFormat = "^([a-z]+-[0-9]+)$"

	// AlphanumericCapsFormat represents alphanumeric caps ticket format e.g. FEAT-123
	AlphanumericCapsFormat IDFormat = "^([A-Z]+-[0-9]+)$"

	// NumericFormat represents numeric ticket format e.g. 123
	NumericFormat IDFormat = "^#?([0-9]+)$"
)

func (i IDFormat) Options() []IDFormat {
	return []IDFormat{
		AlphanumericFormat,
		AlphanumericSmallFormat,
		AlphanumericCapsFormat,
		NumericFormat,
	}
}

// String returns string representation of ticket id
func (t ID) String() string {
	return string(t)
}

// ParseID parses ticket id from string
func ParseID(s string, format IDFormat) (ID, error) {
	re := regexp.MustCompile(string(format))
	match := re.FindStringSubmatch(s)
	if len(match) < 2 || len(match[1]) < 1 {
		return "", ErrInvalidFormat
	}
	return ID(match[1]), nil
}
