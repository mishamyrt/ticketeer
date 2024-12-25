package pattern_test

import (
	"fmt"
	"testing"

	"github.com/mishamyrt/ticketeer/pkg/pattern"
)

func TestWildcardMatch(t *testing.T) {
	var tests = []struct {
		pattern string
		input   string
		want    bool
	}{
		{"", "", true},
		{"feature/*", "feature/foo", true},
		{"feature/*/*", "feature/foo", false},
		{"feature/*/*", "feature/foo/bar", true},
		{"../*", "./foo", false},
		{"../*", "*/foo", false},
		{"../*", "../foo", true},
		{"[/]/*", "*", false},
		{"[/]/*", "[/]/[]", true},
		{"??*??", "!!foo!!", false},
		{"??*??", "**foo**", false},
		{"??*??", "??foo??", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%s", tt.pattern, tt.input), func(t *testing.T) {
			matcher := pattern.NewWildcard(tt.pattern)
			got := matcher.Match(tt.input)
			if got != tt.want {
				t.Errorf("Match() got = %v, want %v", got, tt.want)
			}
		})
	}

	// t.Run("error", func(t *testing.T) {
	// 	var b strings.Builder
	// 	for i := 0; i < 104123; i++ {
	// 		b.WriteString("*")
	// 	}
	// 	_, err := pattern.NewWildcard(b.String())
	// 	if err == nil {
	// 		t.Errorf("Match() error = %v, wantErr %v", err, true)
	// 	}
	// })
}
