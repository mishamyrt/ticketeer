package config

import (
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/pkg/tpl"
)

// TicketConfig represents ticket configuration
type TicketConfig struct {
	Format     ticket.IDFormat
	AllowEmpty bool
}

// BranchConfig represents branch configuration
type BranchConfig struct {
	Format ticket.BranchFormat
	Ignore []string
}

// MessageConfig represents message configuration
type MessageConfig struct {
	Location TicketLocation
	Template tpl.Template
}

// Config represents configuration
type Config struct {
	Ticket  TicketConfig
	Branch  BranchConfig
	Message MessageConfig
}
