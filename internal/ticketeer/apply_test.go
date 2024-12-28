package ticketeer_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/ticketeer"
)

func TestApplyDryRun(t *testing.T) {
	t.Parallel()

	repoDir := t.TempDir()
	_ = os.WriteFile(filepath.Join(repoDir, "file"), []byte{}, 0644)

	repo, _ := git.NewRepository(repoDir)
	_, _ = repo.Exec(git.Command("switch", "-c", "feature/XXX-123"))

	out := strings.Builder{}
	app := ticketeer.New()
	app.SetupLog(ticketeer.LogOptions{
		Output: &out,
	})
	err := app.Apply(repoDir, &ticketeer.ApplyArgs{
		DryRunWith: "commit",
	})
	if err != nil {
		t.Errorf("Apply() error = %v", err)
	}
	if !strings.Contains(out.String(), "commit\n\nXXX-123") {
		t.Errorf("Apply() output = %v, want %v", out.String(), "commit\n\nXXX-123")
	}

}

func TestApplyMissingConfig(t *testing.T) {
	t.Parallel()

	repoDir := t.TempDir()
	out := strings.Builder{}
	app := ticketeer.New()
	app.SetupLog(ticketeer.LogOptions{
		Output: &out,
	})
	err := app.Apply(repoDir, &ticketeer.ApplyArgs{
		ConfigPath: "./unknown.yml",
	})
	if err == nil {
		t.Errorf("Apply() error = %v, wantErr %v", err, true)
	}
}

func TestApplyMissingRepo(t *testing.T) {
	t.Parallel()

	repoDir := t.TempDir()
	out := strings.Builder{}
	app := ticketeer.New()
	app.SetupLog(ticketeer.LogOptions{
		Output: &out,
	})
	err := app.Apply(repoDir, &ticketeer.ApplyArgs{})
	if err == nil {
		t.Errorf("Apply() error = %v, wantErr %v", err, true)
	}
}

func TestDetached(t *testing.T) {
	t.Parallel()

	repoDir := t.TempDir()
	repo, _ := git.NewRepository(repoDir)
	_ = os.WriteFile(filepath.Join(repoDir, "file"), []byte{}, 0644)

	_, _ = repo.Exec(git.Command("add", "file"))
	_, _ = repo.Exec(git.Command("commit", "-m", "Initial commit"))
	_, _ = repo.Exec(git.Command("git", "tag", "v1.0.0"))
	commit, _ := repo.Exec(git.Command("rev-parse", "HEAD"))
	_, _ = repo.Exec(git.Command("checkout", commit))

	out := strings.Builder{}
	app := ticketeer.New()
	app.SetupLog(ticketeer.LogOptions{
		Output: &out,
	})

	err := app.Apply(repoDir, &ticketeer.ApplyArgs{})
	if err != nil {
		t.Errorf("Apply() error = %v", err)
	}
	if !strings.Contains(out.String(), "skipping") {
		t.Errorf("Apply() output = %v, want %v", out.String(), "skipping")
	}
}
