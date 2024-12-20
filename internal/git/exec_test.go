package git_test

import (
	"strings"
	"testing"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/testutil"
)

func TestIsAvailable(t *testing.T) {
	// This test relies on the environment.
	// If git is not available, test will fail.
	if !git.IsAvailable() {
		t.Fatal("expected git to be available")
	}

	pathMock := testutil.NewEnvMock("PATH", "")
	if git.IsAvailable() {
		t.Fatal("expected git to not be available")
	}
	pathMock.Restore()
}

func TestExec(t *testing.T) {
	var tests = []struct {
		repoPath string
		cmd      string
		validate func(*testing.T, string, error)
	}{
		{"../../", "status", func(tt *testing.T, _ string, err error) {
			if err != nil {
				t.Errorf("Exec() error = %v", err)
				return
			}
		}},
		{"", "--version", func(tt *testing.T, _ string, err error) {
			if err != nil {
				t.Errorf("Exec() error = %v", err)
				return
			}
		}},
		{"../non-existent", "status", func(tt *testing.T, _ string, err error) {
			if err == nil {
				t.Errorf("Exec() got = %v, want error", err)
				return
			} else if !strings.Contains(err.Error(), "no such file") {
				t.Errorf("Exec() unexpected error = %v", err)
				return
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.repoPath, func(t *testing.T) {
			cmd := git.Command("status")
			got, err := cmd.ExecuteAt(tt.repoPath)
			tt.validate(t, got, err)
		})
	}
}
