package config

import "github.com/mishamyrt/ticketeer/internal/tpl"

type Config struct {
	AllowEmpty     bool
	TicketLocation TicketLocation
	Template       tpl.Template
}
