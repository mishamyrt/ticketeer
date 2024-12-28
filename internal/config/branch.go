package config

import (
	"errors"
	"fmt"

	"github.com/mishamyrt/ticketeer/internal/ticket"
)

// ErrUnknownBranchFormat is returned when branch format is unknown
var ErrUnknownBranchFormat = errors.New("unknown branch format")

// BranchFormat represents branch format
type BranchFormat string

const (
	// GitFlowBranch represents git-flow branch format e.g. feature/FEAT-123[/description]
	GitFlowBranch BranchFormat = "git-flow"
	// GitFlowTypelessBranch represents git-flow branch format without type e.g. FEAT-123[/description]
	GitFlowTypelessBranch BranchFormat = "git-flow-typeless"
	// TicketIDBranch represents ticket id branch format e.g. FEAT-123
	TicketIDBranch BranchFormat = "ticket-id"
)

// BranchFormatOptions available branch formats
var BranchFormatOptions = [...]BranchFormat{
	GitFlowBranch,
	GitFlowTypelessBranch,
	TicketIDBranch,
}

// String returns string representation of branch format
func (b BranchFormat) String() string {
	return string(b)
}

var branchFormatMapping = map[BranchFormat]ticket.BranchFormat{
	GitFlowBranch:         ticket.GitFlowBranch,
	GitFlowTypelessBranch: ticket.GitFlowTypelessBranch,
	TicketIDBranch:        ticket.TicketIDBranch,
}

// TicketFormat returns ticket format
func (b BranchFormat) TicketFormat() ticket.BranchFormat {
	return branchFormatMapping[b]
}

// ParseBranchFormat parses branch format
func ParseBranchFormat(s string) (BranchFormat, error) {
	format := BranchFormat(s)
	if _, ok := branchFormatMapping[format]; ok {
		return format, nil
	}
	return "", fmt.Errorf("%w: %s", ErrUnknownBranchFormat, s)
}
