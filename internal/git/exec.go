package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ErrCommandFailed is returned when git command execution fails
var ErrCommandFailed = fmt.Errorf("command failed")

// Cmd represents git command
type Cmd struct {
	// Cmd is native git command instance
	*exec.Cmd
	args []string
}

// Execute executes git command
func (c *Cmd) Execute() (string, error) {
	output, err := c.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrCommandFailed, err)
	}
	return strings.Trim(string(output), "\n"), nil
}

// ExecuteAt executes git command at given path
func (c *Cmd) ExecuteAt(path string) (string, error) {
	if path == "" {
		path = "."
	}
	_, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	previousDir := c.Dir
	c.Dir = path
	output, err := c.Execute()
	c.Dir = previousDir
	return output, err
}

// Command returns new git command
func Command(args ...string) *Cmd {
	return &Cmd{
		args: args,
		Cmd:  exec.Command("git", args...),
	}
}

// IsAvailable returns true if git is available in PATH
func IsAvailable() bool {
	_, err := exec.LookPath("git")
	return err == nil
}
