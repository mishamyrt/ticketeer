package format_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/internal/config"
	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/internal/ticketeer/format"
	"github.com/mishamyrt/ticketeer/internal/tpl"
)

func TestTitle(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		template tpl.Template
		title    string
		ticket   ticket.ID
		want     string
		wantErr  bool
	}{
		{
			"{ticket}:",
			"feat: add amazing feature",
			"XXX-123",
			"XXX-123: feat: add amazing feature",
			false,
		},
		{
			"#{ticket}",
			"feat: add amazing feature",
			"XXX-123",
			"#XXX-123 feat: add amazing feature",
			false,
		},
		{
			"{ticket}",
			"feat: add amazing feature",
			"",
			"feat: add amazing feature",
			false,
		},
		{
			"{ticket}:",
			"XXX-123: feat: add amazing feature",
			"XXX-123",
			"XXX-123: feat: add amazing feature",
			false,
		},
		{
			"{ticket}{unknown}",
			"feat: add amazing feature",
			"XXX-123",
			"",
			true,
		},
		{
			"{unknown}",
			"feat: add amazing feature",
			"XXX-123",
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got, err := format.Title(tt.template, tt.title, tt.ticket)
			if (err != nil) != tt.wantErr {
				t.Errorf("Title() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Title() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBody(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		template tpl.Template
		body     string
		ticket   ticket.ID
		want     string
		wantErr  bool
	}{
		{
			"{ticket}",
			"",
			"XXX-123",
			"XXX-123",
			false,
		},
		{
			"Issue: {ticket}",
			"",
			"XXX-123",
			"Issue: XXX-123",
			false,
		},
		{
			"{ticket}",
			"Description: some text",
			"XXX-123",
			"Description: some text\n\nXXX-123",
			false,
		},
		{
			"{ticket}",
			"Description: some text\n\nXXX-123",
			"XXX-123",
			"Description: some text\n\nXXX-123",
			false,
		},
		{
			"{ticket}",
			"XXX-123\n\nDescription: some text",
			"",
			"XXX-123\n\nDescription: some text",
			false,
		},
		{
			"{unknown}",
			"",
			"XXX-123",
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.body, func(t *testing.T) {
			got, err := format.Body(tt.template, tt.body, tt.ticket)
			if (err != nil) != tt.wantErr {
				t.Errorf("Body() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Body() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		template tpl.Template
		message  git.CommitMessage
		location config.TicketLocation
		ticket   ticket.ID
		want     git.CommitMessage
		wantErr  bool
	}{
		{
			"{ticket}:",
			git.CommitMessage{Title: "feat: add amazing feature"},
			config.TicketLocationTitle,
			"XXX-123",
			git.CommitMessage{Title: "XXX-123: feat: add amazing feature"},
			false,
		},
		{
			"{ticket}",
			git.CommitMessage{Title: "feat: add amazing feature"},
			config.TicketLocationBody,
			"XXX-123",
			git.CommitMessage{Title: "feat: add amazing feature", Body: "XXX-123"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.message.String(), func(t *testing.T) {
			got := tt.message
			err := format.Message(&got, tt.ticket, config.MessageConfig{
				Location: tt.location,
				Template: tt.template,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}
