package ticket_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/internal/ticket"
)

func TestParseFromBranchName(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		branchName string
		want       ticket.ID
		format     ticket.Format
	}{
		{"feature/FEAT-123", "FEAT-123", ticket.AlphanumericFormat},
		{"feature/123/text", "123", ticket.NumericFormat},
		{"sprint-12", "", ticket.AlphanumericFormat},
		{"main", "", ticket.AlphanumericFormat},
		{"01104818870da818fe9c43e7b9e3b5946997c175", "", ticket.AlphanumericFormat},
		{"feature/", "", ticket.AlphanumericFormat},
		{"feature/FEAT-123/", "FEAT-123", ticket.AlphanumericFormat},
		{"feature//", "", ticket.AlphanumericFormat},
	}

	for _, tt := range tests {
		t.Run(tt.branchName, func(t *testing.T) {
			got, err := ticket.ParseFromBranchName(tt.branchName, tt.format)
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
