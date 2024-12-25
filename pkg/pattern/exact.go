package pattern

// Exact represents an exact value matcher
type Exact struct {
	value string
}

// NewExact creates a new exact matcher
func NewExact(s string) Matcher {
	return &Exact{value: s}
}

// Match returns true if the string matches the exact value
func (m *Exact) Match(s string) bool {
	return s == m.value
}
