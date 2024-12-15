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
	return template.Render(tpl.Variables{
		"ticket": id.String(),
		"title":  title,
	})
}

// Body renders the body of the message
func Body(
	template tpl.Template,
	body string,
	id ticket.ID,
) (string, error) {
	return template.Render(tpl.Variables{
		"ticket": id.String(),
		"body":   body,
	})
}

// Message appends ticket id to commit message
func Message(message *git.CommitMessage, ticketID ticket.ID, cfg config.Config) error {
	var err error
	switch cfg.TicketLocation {
	case config.TicketLocationTitle:
		message.Title, err = Title(cfg.Template, message.Title, ticketID)
	case config.TicketLocationBody:
		message.Body, err = Body(cfg.Template, message.Body, ticketID)
	}
	return err
}
