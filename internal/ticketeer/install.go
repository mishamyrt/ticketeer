package ticketeer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/hook"
	"github.com/mishamyrt/ticketeer/pkg/log"
	"github.com/mishamyrt/ticketeer/pkg/log/color"
)

const readmeURL = "https://github.com/mishamyrt/ticketeer?tab=readme-ov-file"

// Install git hook
func (a *App) Install(force bool) error {
	repo, err := git.OpenRepository("./")
	if err != nil {
		return err
	}
	log.Debugf("Repository root found at: %s", repo.Path())
	hookPath, err := getHookPath(repo)
	if err != nil {
		return err
	}

	// Install hook if it does not exist
	_, err = os.Stat(hookPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug("Hook not found")
			return a.installHook(hookPath)
		}
		return err
	}

	if force {
		log.Debug("Force installing hook")
		return a.installHook(hookPath)
	}

	// Check installed hook runner
	runner, err := hook.DetectRunner(hookPath)
	if err != nil {
		if !errors.Is(err, hook.ErrUnknownRunner) {
			return err
		}
		log.Info(color.Yellow("Detected unknown hook."))
		log.Info("To replace the hook, run:")
		log.Info(color.Cyan("  ticketeer install --force"))
		return nil
	}

	if runner.GuideAnchor == "" {
		log.Info("Hook already installed")
		return nil
	}

	log.Infof(color.Yellow(
		"Detected %s. You can use it in tandem with ticketeer!",
	), runner.Name)
	setupURL := fmt.Sprintf("%s#%s", readmeURL, runner.GuideAnchor)
	log.Infof("Setup instructions: %s\n", color.Yellow(setupURL))
	log.Info("To replace the hook, run:")
	log.Info(color.Cyan("  ticketeer install --force"))

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
	log.Info(color.Green("Hook successfully installed"))
	return nil
}

func getHookPath(repo *git.Repository) (string, error) {
	hooksDir, err := repo.HooksDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(hooksDir, hook.Name), nil
}
