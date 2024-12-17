package ticketeer

import (
	"errors"
	"fmt"
	"os"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/hook"
)

const readmeURL = "https://github.com/mishamyrt/ticketeer?tab=readme-ov-file"

// Install git hook
func Install(_ *Options, force bool) error {
	repo, err := git.OpenRepository("./")
	if err != nil {
		return err
	}
	hookPath, err := getHookPath(repo)
	if err != nil {
		return err
	}

	// Install hook if it does not exist
	_, err = os.Stat(hookPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Installing hook...")
			return installHook(hookPath)
		}
		return err
	}

	if force {
		fmt.Println("Replacing hook...")
		return installHook(hookPath)
	}

	// Check installed hook runner
	runner, err := hook.DetectRunner(hookPath)
	if err != nil {
		if errors.Is(err, hook.ErrUnknownRunner) {
			fmt.Println("Detected unknown hook.")
			fmt.Println("To replace the hook, run:")
			fmt.Println("  ticketeer install --force")
		}
		return err
	}

	fmt.Printf("Detected %s. You can use it in tandem with ticketeer!\n", runner.Name)
	fmt.Printf("Setup instructions: %s#%s\n", readmeURL, runner.GuideAnchor)
	fmt.Println("To replace the hook, run:")
	fmt.Println("  ticketeer install --force")

	return err
}

func installHook(hookPath string) error {
	fmt.Println("⚙️ Installing hook...")
	err := os.WriteFile(hookPath, hook.Content(), 0755)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	fmt.Println("🚀 Hook installed")
	return nil
}

func getHookPath(repo git.Repository) (string, error) {
	hooksDir, err := repo.HooksPath()
	if err != nil {
		return "", err
	}
	return hook.Path("prepare-commit-msg", hooksDir), nil
}
