package config

import (
	"errors"
	"fmt"
)

// ErrUnknownLocation is returned when ticket location is unknown
var ErrUnknownLocation = errors.New("unknown ticket location")

// TicketLocation represents ticket location
type TicketLocation string

const (
	// TicketLocationTitle represents ticket id located in title
	TicketLocationTitle TicketLocation = "title"

	// TicketLocationBody represents ticket id located in body
	TicketLocationBody TicketLocation = "body"
)

// ParseLocation parses ticket location from string
func ParseLocation(location string) (TicketLocation, error) {
	switch TicketLocation(location) {
	case TicketLocationTitle:
		return TicketLocationTitle, nil
	case TicketLocationBody:
		return TicketLocationBody, nil
	}
	return "", fmt.Errorf("%w: %s", ErrUnknownLocation, location)
}
