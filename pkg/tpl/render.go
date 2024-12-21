package tpl

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	// ErrUndefinedVariables is returned when template contains undefined variables
	ErrUndefinedVariables = errors.New("undefined variables")

	variableRe = regexp.MustCompile(`{(.+?)}`)
)

const undefinedKeyword = "undefined"

func newErrorUndefinedVariables(variables []string) error {
	return fmt.Errorf("%w: %s", ErrUndefinedVariables, strings.Join(variables, ", "))
}

// Render the template with the given variables
func (t Template) Render(variables Variables) (string, error) {
	undefinedVariables := make([]string, 0)
	result := variableRe.ReplaceAllStringFunc(string(t), func(match string) string {
		key := match[1 : len(match)-1]
		if value, ok := variables[key]; ok {
			return value
		}
		undefinedVariables = append(undefinedVariables, key)
		return undefinedKeyword
	})

	if len(undefinedVariables) > 0 {
		return result, newErrorUndefinedVariables(undefinedVariables)
	}
	return result, nil
}
