package ticketeer_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/hook"
	"github.com/mishamyrt/ticketeer/internal/ticketeer"
)

func TestInstall(t *testing.T) {
	t.Parallel()

	app := ticketeer.New()
	out := strings.Builder{}
	app.SetupLog(ticketeer.LogOptions{
		Output: &out,
	})

	repoDir := t.TempDir()
	_, err := git.NewRepository(repoDir)
	if err != nil {
		t.Fatalf("NewRepository() error = %v", err)
	}
	err = app.Install(repoDir, false)
	if err != nil {
		t.Errorf("Install() error = %v", err)
	}
}

func TestInstallDetectRepo(t *testing.T) {
	t.Parallel()

	app := ticketeer.New()
	out := strings.Builder{}
	app.SetupLog(ticketeer.LogOptions{
		Output: &out,
	})

	repoDir := t.TempDir()

	err := app.Install(repoDir, false)
	if err == nil {
		t.Errorf("Install() error = %v, wantErr %v", err, true)
	}
}

func TestInstallDetectRunner(t *testing.T) {
	t.Parallel()

	app := ticketeer.New()
	out := strings.Builder{}
	app.SetupLog(ticketeer.LogOptions{
		Output: &out,
	})

	repoDir := t.TempDir()
	repo, err := git.NewRepository(repoDir)
	if err != nil {
		t.Fatalf("NewRepository() error = %v", err)
	}
	err = app.Install(repoDir, false)
	if err != nil {
		t.Errorf("Install() error = %v", err)
	}

	err = app.Install(repoDir, false)
	if err == nil {
		t.Errorf("Install() error = %v, wantErr %v", err, true)
	}

	err = copyHook("unknown", repo)
	if err != nil {
		t.Fatalf("copyHook() error = %v", err)
	}

	err = app.Install(repoDir, false)
	if err == nil {
		t.Errorf("Install() error = %v, wantErr %v", err, true)
	}
	if !strings.Contains(out.String(), "unknown") {
		t.Errorf("Install() output = %s, want %s", out.String(), "unknown")
	}

	out.Reset()
	err = copyHook("lefthook", repo)
	if err != nil {
		t.Fatalf("copyHook() error = %v", err)
	}
	err = app.Install(repoDir, false)
	if err == nil {
		t.Errorf("Install() error = %v, wantErr %v", err, true)
	}
	if !strings.Contains(out.String(), "lefthook") {
		t.Errorf("Install() output = %s, want %s", out.String(), "lefthook")
	}
}

func copyHook(name string, repo *git.Repository) error {
	srcPath := filepath.Join("..", "..", "testdata", "hook_runners", name+".sh")
	dstPath := filepath.Join(repo.HooksDir(), hook.Name)
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	_, err = os.Stat(dstPath)
	if !os.IsNotExist(err) {
		err = os.Remove(dstPath)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(dstPath, content, 0755)
}
