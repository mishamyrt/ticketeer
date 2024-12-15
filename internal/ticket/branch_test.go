package ticket_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/internal/ticket"
)

func TestFindInBranch(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		branchName string
		want       string
		format     ticket.BranchFormat
	}{
		{"feature/FEAT-123", "FEAT-123", ticket.GitFlowBranch},
		{"feature/123/text", "123", ticket.GitFlowBranch},
		{"feature/FEAT-123/", "FEAT-123", ticket.GitFlowBranch},
		{"feature//", "", ticket.GitFlowBranch},
		{"sprint-12", "sprint-12", ticket.TicketIDBranch},
		{"sprint-12", "", ticket.GitFlowBranch},
		{"main", "", ticket.GitFlowBranch},
		{"01104818870da818fe9c43e7b9e3b5946997c175", "", ticket.GitFlowBranch},
		{"FEAT-123/description", "FEAT-123", ticket.GitFlowTypelessBranch},
		{"FEAT-123", "FEAT-123", ticket.GitFlowTypelessBranch},
		{"FEAT-123", "FEAT-123", ticket.TicketIDBranch},
	}

	for _, tt := range tests {
		t.Run(tt.branchName, func(t *testing.T) {
			got, err := ticket.FindInBranch(tt.branchName, tt.format)
			if (err != nil) != (len(tt.want) == 0) {
				t.Errorf("ParseFromBranchName() case %s, error = %v, want %v", tt.branchName, err, tt.want)
				return
			}
			if got != tt.want {
				t.Errorf("ParseFromBranchName() case %s got = %v, want %v", tt.branchName, got, tt.want)
			}
		})
	}
}
