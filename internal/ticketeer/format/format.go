package format

import (
	"strings"

	"github.com/mishamyrt/ticketeer/internal/config"
	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/pkg/tpl"
)

// Title renders the title of the message
func Title(
	template tpl.Template,
	title string,
	id ticket.ID,
) (string, error) {
	ticketPrefix, err := template.Render(tpl.Variables{
		"ticket": id.String(),
	})
	if err != nil {
		return "", err
	}
	// Check if the title already contains the ticket id
	if strings.HasPrefix(title, ticketPrefix) {
		return title, nil
	}
	return ticketPrefix + " " + title, nil
}

// Body renders the body of the message
func Body(
	template tpl.Template,
	body string,
	id ticket.ID,
) (string, error) {
	ticketLine, err := template.Render(tpl.Variables{
		"ticket": id.String(),
	})
	if err != nil {
		return "", err
	}
	if body == "" {
		return ticketLine, nil
	}
	// Check if the body already contains the ticket id
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if strings.TrimRight(line, "\n") == ticketLine {
			return body, nil
		}
	}
	return body + "\n\n" + ticketLine, nil
}

// Message appends ticket id to commit message
func Message(message *git.CommitMessage, ticketID ticket.ID, cfg config.MessageConfig) error {
	var err error
	switch cfg.Location {
	case config.TicketLocationTitle:
		message.Title, err = Title(cfg.Template, message.Title, ticketID)
	case config.TicketLocationBody:
		message.Body, err = Body(cfg.Template, message.Body, ticketID)
	}
	return err
}
