package config

import (
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/pkg/tpl"
)

var defaultTemplates = map[TicketLocation]tpl.Template{
	TicketLocationTitle: "{ticket}:",
	TicketLocationBody:  "{ticket}",
}

// Default is the default configuration
var Default = Config{
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

// DefaultFileName is the default file name of the configuration file
const DefaultFileName = ".ticketeer.yaml"
