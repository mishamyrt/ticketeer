package ticket_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/internal/ticket"
)

func TestAssert(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		id      string
		want    ticket.ID
		format  ticket.IDFormat
		wantErr bool
	}{
		{"FEAT-123", "FEAT-123", ticket.AlphanumericFormat, false},
		{"feat-1", "feat-1", ticket.AlphanumericFormat, false},
		{"a-1", "a-1", ticket.AlphanumericFormat, false},
		{"a", "", ticket.AlphanumericFormat, true},
		{"1", "", ticket.AlphanumericFormat, true},
		{"", "", ticket.AlphanumericFormat, true},

		{"FEAT-123", "", ticket.NumericFormat, true},
		{"feat-1", "", ticket.NumericFormat, true},
		{"a-1", "", ticket.NumericFormat, true},
		{"a", "", ticket.NumericFormat, true},
		{"1", "1", ticket.NumericFormat, false},
		{"123", "123", ticket.NumericFormat, false},
		{"0", "0", ticket.NumericFormat, false},
		{"0000", "0000", ticket.NumericFormat, false},
		{"", "", ticket.NumericFormat, true},
		{"#123", "123", ticket.NumericFormat, false},
		{"#0", "0", ticket.NumericFormat, false},
		{"#", "", ticket.NumericFormat, true},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got, err := ticket.ParseID(tt.id, tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("Assert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Assert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
