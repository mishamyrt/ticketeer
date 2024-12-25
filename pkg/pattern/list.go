package pattern

import "strings"

// Matcher is a pattern matcher interface
type Matcher interface {
	Match(path string) bool
}

// List is a special matcher that uses a list of patterns
type List struct {
	patterns []Matcher
}

// NewList creates a new list matcher
func NewList(patterns ...string) Matcher {
	p := make([]Matcher, 0, len(patterns))
	for _, pattern := range patterns {
		if strings.Contains(pattern, "*") {
			p = append(p, NewWildcard(pattern))
			continue
		}
		p = append(p, NewExact(pattern))
	}
	return &List{patterns: p}
}

// Match matches a path against a list of patterns
func (l *List) Match(path string) bool {
	for _, pattern := range l.patterns {
		if pattern.Match(path) {
			return true
		}
	}
	return false
}
