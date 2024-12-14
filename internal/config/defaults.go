package config

import "github.com/mishamyrt/ticketeer/internal/tpl"

var defaultTemplates = map[TicketLocation]tpl.Template{
	TicketLocationTitle: "{ticket}: {title}",
	TicketLocationBody:  "{body}\n\n{ticket}",
}

var defaultConfig = Config{
	AllowEmpty:     true,
	TicketLocation: TicketLocationBody,
	Template:       defaultTemplates[TicketLocationTitle],
}

const DefaultPath = "./ticketeer.yaml"
