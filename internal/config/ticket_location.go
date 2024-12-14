package config

type TicketLocation string

const (
	TicketLocationTitle TicketLocation = "title"
	TicketLocationBody  TicketLocation = "body"
)

func ParseLocation(location string) (TicketLocation, error) {
	switch TicketLocation(location) {
	case TicketLocationTitle:
		return TicketLocationTitle, nil
	case TicketLocationBody:
		return TicketLocationBody, nil
	}
	return "", ErrUnknownLocation
}
