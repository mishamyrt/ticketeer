package git

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// ErrBranchNotFound is returned when repository head is in detached state
var ErrBranchNotFound = errors.New("branch is not found")

// ErrHeadNotFound is returned when head file is not found
var ErrHeadNotFound = errors.New("head file is not found")

const headPrefix = "ref: refs/heads/"

var headPath = filepath.Join(".git", "HEAD")

// BranchName represents git branch name
type BranchName string

// String returns string representation of branch name
func (b BranchName) String() string {
	return string(b)
}

// ParseBranchName parses branch name from head content
func ParseBranchName(head string) (BranchName, error) {
	if !strings.HasPrefix(head, headPrefix) {
		return "", ErrBranchNotFound
	}
	head = strings.TrimPrefix(head, headPrefix)
	head = strings.TrimRight(head, " \n")
	return BranchName(head), nil
}

// ReadBranchName reads branch name from internal git file.
func ReadBranchName() (BranchName, error) {
	rawContent, err := os.ReadFile(headPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = ErrHeadNotFound
		}
		return "", err
	}
	return ParseBranchName(string(rawContent))
}
