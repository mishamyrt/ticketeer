package color

import (
	"fmt"

	"github.com/fatih/color"
)

// ColorizeFunc is a function that colorizes text.
type ColorizeFunc func(v ...any) string

var noColor bool = color.NoColor

// SetNoColor sets if the output should be colorized or not.
func SetNoColor(enabled bool) {
	noColor = enabled
	color.NoColor = noColor
}

// WrapNoColor returns a function that wraps the given colorize function and
// returns the given text when the output is not colorized.
func WrapNoColor(colorize ColorizeFunc) ColorizeFunc {
	return func(a ...interface{}) string {
		if noColor {
			return fmt.Sprint(a...)
		}
		return colorize(a...)
	}
}
