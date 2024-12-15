package format

import (
	"github.com/mishamyrt/ticketeer/internal/config"
	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/internal/tpl"
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
