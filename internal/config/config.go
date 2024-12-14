package config

import "github.com/mishamyrt/ticketeer/internal/tpl"

// Config represents configuration
type Config struct {
	AllowEmpty     bool
	TicketLocation TicketLocation
	Template       tpl.Template
}
