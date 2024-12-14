package render

import (
	"github.com/mishamyrt/ticketeer/internal/ticket"
	"github.com/mishamyrt/ticketeer/internal/tpl"
)

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
