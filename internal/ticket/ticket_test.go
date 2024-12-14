package ticket_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/internal/ticket"
)

func TestAssert(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		id      string
		format  ticket.Format
		wantErr bool
	}{
		{"FEAT-123", ticket.AlphanumericFormat, false},
		{"feat-1", ticket.AlphanumericFormat, false},
		{"a-1", ticket.AlphanumericFormat, false},
		{"a", ticket.AlphanumericFormat, true},
		{"1", ticket.AlphanumericFormat, true},
		{"", ticket.AlphanumericFormat, true},

		{"FEAT-123", ticket.NumericFormat, true},
		{"feat-1", ticket.NumericFormat, true},
		{"a-1", ticket.NumericFormat, true},
		{"a", ticket.NumericFormat, true},
		{"1", ticket.NumericFormat, false},
		{"123", ticket.NumericFormat, false},
		{"0", ticket.NumericFormat, false},
		{"0000", ticket.NumericFormat, false},
		{"", ticket.NumericFormat, true},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			if _, err := ticket.ParseID(tt.id, tt.format); (err != nil) != tt.wantErr {
				t.Errorf("Assert() case %s error = %v, wantErr %v", tt.id, err, tt.wantErr)
			}
		})
	}
}
