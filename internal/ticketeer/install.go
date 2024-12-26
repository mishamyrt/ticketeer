package ticketeer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/hook"
	"github.com/mishamyrt/ticketeer/pkg/log/color"
)

const readmeURL = "https://github.com/mishamyrt/ticketeer?tab=readme-ov-file"

// Install git hook
func (a *App) Install(workingDir string, force bool) error {
	repo, err := git.OpenRepository(workingDir)
	if err != nil {
		return a.handleError(err, "Failed to open repository")
	}
	a.log.Debugf("Repository root found at: %s", repo.Path())
	hookPath, err := getHookPath(repo)
	if err != nil {
		return a.handleError(err, "Failed to get hook path")
	}

	// Install hook if it does not exist
	_, err = os.Stat(hookPath)
	if err != nil {
		if os.IsNotExist(err) {
			a.log.Debug("Hook not found")
			return a.installHook(hookPath)
		}
		return a.handleError(err, "Failed to get hook stat")
	}

	if force {
		a.log.Debug("Force installing hook")
		return a.installHook(hookPath)
	}

	// Check installed hook runner
	runner, err := hook.DetectRunner(hookPath)
	if err != nil {
		if !errors.Is(err, hook.ErrUnknownRunner) {
			return a.handleError(err, "Failed to detect hook runner")
		}
		a.log.Info(color.Yellow("Detected unknown hook."))
		a.log.Info("To replace the hook, run:")
		a.log.Info(color.Cyan("  ticketeer install --force"))
		return nil
	}

	if runner.GuideAnchor == "" {
		a.log.Info("Hook already installed")
		return nil
	}

	a.log.Infof(color.Yellow(
		"Detected %s. You can use it in tandem with ticketeer!",
	), runner.Name)
	setupURL := fmt.Sprintf("%s#%s", readmeURL, runner.GuideAnchor)
	a.log.Infof("Setup instructions: %s\n", color.Yellow(setupURL))
	a.log.Info("To replace the hook, run:")
	a.log.Info(color.Cyan("  ticketeer install --force"))

	return nil
}

func (a *App) installHook(hookPath string) error {
	content, err := hook.Content()
	if err != nil {
		return err
	}
	err = os.WriteFile(hookPath, content, 0755)
	if err != nil {
		return a.handleError(err, "Failed to write hook")
	}
	a.log.Info(color.Green("Hook successfully installed"))
	return nil
}

func getHookPath(repo *git.Repository) (string, error) {
	hooksDir, err := repo.HooksDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(hooksDir, hook.Name), nil
}
