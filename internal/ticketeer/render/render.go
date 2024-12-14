package render

import (
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
