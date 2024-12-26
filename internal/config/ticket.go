package config

import (
	"errors"
	"fmt"

	"github.com/mishamyrt/ticketeer/internal/ticket"
)

var (
	// ErrUnknownLocation is returned when ticket location is unknown
	ErrUnknownLocation = errors.New("unknown ticket location")

	// ErrUnknownTicketFormat is returned when ticket format is unknown
	ErrUnknownTicketFormat = errors.New("unknown ticket format")
)

// TicketLocation represents ticket location
type TicketLocation string

const (
	// TicketLocationTitle represents ticket id located in title
	TicketLocationTitle TicketLocation = "title"

	// TicketLocationBody represents ticket id located in body
	TicketLocationBody TicketLocation = "body"
)

func (ticket TicketLocation) Options() []TicketLocation {
	return []TicketLocation{
		TicketLocationTitle,
		TicketLocationBody,
	}
}

// ParseTicketLocation parses ticket location from string
func ParseTicketLocation(location string) (TicketLocation, error) {
	switch TicketLocation(location) {
	case TicketLocationTitle:
		return TicketLocationTitle, nil
	case TicketLocationBody:
		return TicketLocationBody, nil
	}
	return "", fmt.Errorf("%w: %s", ErrUnknownLocation, location)
}

// ParseTicketFormat parses ticket format from string
func ParseTicketFormat(format string) (ticket.IDFormat, error) {
	switch format {
	case "alphanumeric":
		return ticket.AlphanumericFormat, nil
	case "alphanumeric-small":
		return ticket.AlphanumericSmallFormat, nil
	case "alphanumeric-caps":
		return ticket.AlphanumericCapsFormat, nil
	case "numeric":
		return ticket.NumericFormat, nil
	default:
		return "", fmt.Errorf("%w: %s", ErrUnknownTicketFormat, format)
	}
}
