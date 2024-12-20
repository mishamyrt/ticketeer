package git_test

import (
	"slices"
	"testing"

	"github.com/mishamyrt/ticketeer/internal/git"
)

func TestParseCommitMessage(t *testing.T) {
	var tests = []struct {
		message string
		want    git.CommitMessage
	}{
		{"", git.CommitMessage{}},
		{"title", git.CommitMessage{Title: "title"}},
		{"title\n\nbody", git.CommitMessage{Title: "title", Body: "body"}},
		{"title\nbody", git.CommitMessage{Title: "title\nbody"}},
	}

	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			got, err := git.ParseCommitMessage(tt.message)
			if tt.want.Title == "" {
				if err == nil {
					t.Errorf("Parse() got = %v, want error", got)
					return
				}
			} else if err != nil {
				t.Errorf("Parse() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommitMessageFormat(t *testing.T) {
	var tests = []struct {
		message git.CommitMessage
		want    string
	}{
		{git.CommitMessage{}, ""},
		{git.CommitMessage{Title: "title"}, "title"},
		{git.CommitMessage{Title: "title", Body: "body"}, "title\n\nbody"},
		{git.CommitMessage{Title: "title\nbody"}, "title\nbody"},
	}

	for _, tt := range tests {
		t.Run(tt.message.String(), func(t *testing.T) {
			if got := tt.message.String(); got != tt.want {
				t.Errorf("Message.String() got = %v, want %v", got, tt.want)
			}

			gotBytes := tt.message.Bytes()
			if !slices.Equal(gotBytes, []byte(tt.want)) {
				t.Errorf("Message.Bytes() got = %v, want %v", gotBytes, tt.want)
			}
		})
	}
}
