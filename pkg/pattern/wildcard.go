package pattern

import (
	"fmt"
	"regexp"
	"strings"
)

// Wildcard represents a wildcard pattern matcher.
type Wildcard struct {
	re *regexp.Regexp
}

var forbiddenCharsRe = regexp.MustCompile(`[\\\^\$\|\?\+\-\[\]\{\}\(\)]+`)

// NewWildcard creates a new wildcard matcher
func NewWildcard(s string) Matcher {
	expr := forbiddenCharsRe.ReplaceAllStringFunc(s, func(match string) string {
		var b strings.Builder
		for i := 0; i < len(match); i++ {
			b.WriteRune('\\')
			b.WriteByte(match[i])
		}
		return b.String()
	})
	expr = strings.ReplaceAll(expr, "*", ".*")
	expr = fmt.Sprintf("^%s$", expr)
	re := regexp.MustCompile(expr)
	return &Wildcard{re: re}
}

// Match returns true if the string matches the wildcard pattern
func (m *Wildcard) Match(s string) bool {
	return m.re.MatchString(s)
}
