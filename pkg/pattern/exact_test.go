package pattern_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/pkg/pattern"
)

func TestExactMatch(t *testing.T) {
	var tests = []struct {
		pattern string
		input   string
		want    bool
	}{
		{"", "", true},
		{"", "foo", false},
		{"foo", "foo", true},
		{"foo", "bar", false},
		{"foo", "foobar", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			matcher := pattern.NewExact(tt.pattern)
			got := matcher.Match(tt.input)
			if got != tt.want {
				t.Errorf("Match() got = %v, want %v", got, tt.want)
			}
		})
	}
}
