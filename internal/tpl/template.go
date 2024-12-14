package tpl

type Template string

const EmptyTemplate Template = ""

var _ TemplateRenderer = EmptyTemplate

type Variables map[string]string

type TemplateRenderer interface {
	// String returns the template as a string
	String() string

	// Render the template with the given variables
	Render(vars Variables) (string, error)
}

func (t Template) String() string {
	return string(t)
}
