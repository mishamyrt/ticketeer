package color

import (
	"github.com/fatih/color"
)

var (
	// Red is a helper function to print text with red foreground.
	Red = WrapNoColor(color.New(color.FgRed).SprintFunc())

	// Yellow is a helper function to print text with yellow foreground.
	Yellow = WrapNoColor(color.New(color.FgYellow).SprintFunc())

	// Cyan is a helper function to print text with cyan foreground.
	Cyan = WrapNoColor(color.New(color.FgCyan).SprintFunc())

	// Green is a helper function to print text with green foreground.
	Green = WrapNoColor(color.New(color.FgGreen).SprintFunc())

	// Dim is a helper function to print text with dim foreground.
	Dim = WrapNoColor(color.New(color.Faint).SprintFunc())
)
