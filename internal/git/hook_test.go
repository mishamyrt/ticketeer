package git_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/internal/git"
)

func TestHooksPath(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		repoPath string
		want     string
		wantErr  bool
	}{
		// Current repo
		{"../../", ".git/hooks", false},
		// Non-existent repo
		{"../../non-existent", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.repoPath, func(t *testing.T) {
			repo := git.NewRepository(tt.repoPath)
			got, err := repo.HooksPath()
			if (err != nil) != tt.wantErr {
				t.Errorf("HooksPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HooksPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
