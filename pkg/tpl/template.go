package tpl

// Template that can be rendered
type Template string

// EmptyTemplate is an empty template
const EmptyTemplate Template = ""

var _ TemplateRenderer = EmptyTemplate

// Variables that can be used in the template
type Variables map[string]string

// TemplateRenderer is a template renderer interface
type TemplateRenderer interface {
	// String returns the template as a string
	String() string

	// Render the template with the given variables
	Render(vars Variables) (string, error)
}

func (t Template) String() string {
	return string(t)
}
