package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ErrNotRepository is returned when path is not a git repository
var ErrNotRepository = fmt.Errorf("git repository is not found")

// ErrCommandFailed is returned when git command execution fails
var ErrCommandFailed = fmt.Errorf("command failed")

func newErrNotRepository(path string) error {
	return fmt.Errorf("%w: %s", ErrNotRepository, path)
}

// IsAvailable returns true if git is available in PATH
func IsAvailable() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

// IsRepository returns true if path is a git repository
func IsRepository(path string) bool {
	_, err := os.Stat(filepath.Join(path, ".git"))
	return err == nil
}

// AssertRepository returns error if path is not a git repository
func AssertRepository(path string) error {
	if !IsRepository(path) {
		return newErrNotRepository(path)
	}
	return nil
}

// Exec executes git command and returns output
func Exec(path string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, output)
	}
	return string(output), nil
}
