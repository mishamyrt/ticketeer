package pattern_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/pkg/pattern"
)

func TestListMatch(t *testing.T) {
	var tests = []struct {
		patterns []string
		input    string
		want     bool
	}{
		{[]string{"foo"}, "foo", true},
		{[]string{"foo"}, "bar", false},
		{[]string{"foo", "bar"}, "foo", true},
		{[]string{"foo", "bar"}, "bar", true},
		{[]string{"foo", "bar"}, "baz", false},
		{[]string{"foo", "foo/*"}, "foo/bar", true},
		{[]string{"main", "release/*", "develop"}, "dev", false},
		{[]string{"main", "release/*", "develop"}, "release/1.0.0", true},
		{[]string{"main", "release/*", "develop"}, "develop", true},
		{[]string{"main", "release/*", "develop"}, "master", false},
		{[]string{"main", "release/*", "develop"}, "main", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			matcher := pattern.NewList(tt.patterns...)
			got := matcher.Match(tt.input)
			if got != tt.want {
				t.Errorf("Match() got = %v, want %v", got, tt.want)
			}
		})
	}
}
