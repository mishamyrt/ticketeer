package git_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/internal/git"
)

func TestBranchNameFromHead(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		head    string
		want    string
		wantErr bool
	}{
		{"ref: refs/heads/main", "main", false},
		{"ref: refs/heads/feature/#123/description", "feature/#123/description", false},
		{"8c5958a42c6cea124100c15fd23cbf131bdd7621", "", true},
		{"", "", true},
		{"v1.0.0", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.head, func(t *testing.T) {
			got, err := git.BranchNameFromHead(tt.head)
			if (err != nil) != tt.wantErr {
				t.Errorf("BranchNameFromHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BranchNameFromHead() got = %v, want %v", got, tt.want)
			}
		})
	}
}
