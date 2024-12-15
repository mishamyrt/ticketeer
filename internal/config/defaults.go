package config

import (
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/internal/tpl"
)

var defaultTemplates = map[TicketLocation]tpl.Template{
	TicketLocationTitle: "{ticket}:",
	TicketLocationBody:  "{ticket}",
}

var defaultConfig = Config{
	Ticket: TicketConfig{
		Format:     ticket.AlphanumericCapsFormat,
		AllowEmpty: true,
	},
	Branch: BranchConfig{
		Format: ticket.GitFlowBranch,
		Ignore: []string{
			"main",
			"master",
			"develop",
			"dev",
			"release/*",
		},
	},
	Message: MessageConfig{
		Location: TicketLocationBody,
		Template: defaultTemplates[TicketLocationBody],
	},
}

// DefaultPath is the default path to the configuration file
const DefaultPath = "./.ticketeer.yaml"
