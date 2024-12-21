package ticketeer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/hook"
)

const readmeURL = "https://github.com/mishamyrt/ticketeer?tab=readme-ov-file"

// Install git hook
func (a *App) Install(force bool) error {
	repo, err := git.OpenRepository("./")
	if err != nil {
		return err
	}
	a.log.Debugf("Repository root found at: %s", repo.Path())
	hookPath, err := getHookPath(repo)
	if err != nil {
		return err
	}

	// Install hook if it does not exist
	_, err = os.Stat(hookPath)
	if err != nil {
		if os.IsNotExist(err) {
			a.log.Debug("Hook not found")
			return a.installHook(hookPath)
		}
		return err
	}

	if force {
		a.log.Debug("Force installing hook")
		return a.installHook(hookPath)
	}

	// Check installed hook runner
	runner, err := hook.DetectRunner(hookPath)
	if err != nil {
		if !errors.Is(err, hook.ErrUnknownRunner) {
			return err
		}
		a.log.Warn("Detected unknown hook.")
		a.log.Print("To replace the hook, run:")
		a.log.Info("  ticketeer install --force")
		return nil
	}

	if runner.GuideAnchor == "" {
		a.log.Print("Hook already installed")
		return nil
	}

	a.log.Warnf("Detected %s. You can use it in tandem with ticketeer!", runner.Name)
	a.log.Printf("Setup instructions: %s#%s\n", readmeURL, runner.GuideAnchor)
	a.log.Print("To replace the hook, run:")
	a.log.Info("  ticketeer install --force")

	return nil
}

func (a *App) installHook(hookPath string) error {
	content, err := hook.Content()
	if err != nil {
		return err
	}
	err = os.WriteFile(hookPath, content, 0755)
	if err != nil {
		return fmt.Errorf("%s: %w", "Failed to install hook", err)
	}
	a.log.Success("🚀 Hook successfully installed")
	return nil
}

func getHookPath(repo *git.Repository) (string, error) {
	hooksDir, err := repo.HooksDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(hooksDir, hook.Name), nil
}
