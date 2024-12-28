package ticketeer

import (
	"os"
	"path/filepath"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/hook"
	"github.com/mishamyrt/ticketeer/pkg/log/color"
)

// Uninstall git hook
func (a *App) Uninstall(workingDir string, force bool) error {
	repo, err := git.OpenRepository(workingDir)
	if err != nil {
		return a.handleError(err, "Failed to open repository")
	}
	a.log.Debugf("Repository root found at: %s", repo.Path())

	hookPath := filepath.Join(repo.HooksDir(), hook.Name)
	_, err = os.Stat(hookPath)
	if err != nil {
		if os.IsNotExist(err) {
			a.log.Info(color.Red("There is no hook to remove"))
			return wrapHandledError(err)
		}
		return a.handleError(err, "Failed to get hook stat")
	}

	if force {
		a.log.Debug("Force uninstalling hook")
		return a.uninstallHook(hookPath)
	}

	runner, err := hook.DetectRunner(hookPath)
	if err != nil || runner == nil || runner.GuideAnchor != "" {
		a.log.Info(color.Yellow("Detected third party hook."))
		a.log.Info("To uninstall the hook, run:")
		a.log.Info(color.Cyan("  ticketeer uninstall --force"))
		return wrapHandledError(err)
	}

	return a.uninstallHook(hookPath)
}

func (a *App) uninstallHook(hookPath string) error {
	err := os.Remove(hookPath)
	if err != nil {
		return a.handleError(err, "Failed to remove hook")
	}
	a.log.Info(color.Green("Hook successfully uninstalled"))
	return nil
}
