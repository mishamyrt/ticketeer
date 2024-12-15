package ticketeer

import (
	"fmt"
	"os"

	"github.com/mishamyrt/ticketeer/internal/git"
	"github.com/mishamyrt/ticketeer/internal/hook"
)

// Install git hook
func Install(opts *Options, force bool) error {
	if opts.Verbose {
		fmt.Println("VERBOSE")
	}
	repo, err := git.OpenRepository("./")
	if err != nil {
		return err
	}

	hookPath, err := getHookPath(repo)
	if err != nil {
		return err
	}

	installed := hook.FindType(hookPath)

	switch {
	case installed == nil:
		err = installHook(hookPath)
	case force:
		fmt.Println("🚀 Hook is already installed, reinstalling...")
		err = installHook(hookPath)
	default:
		fmt.Printf("Detected %s hook\n", installed.Name())
		fmt.Println(installed.UsageGuide())
	}

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
	return hook.Path(hooksDir), nil
}
