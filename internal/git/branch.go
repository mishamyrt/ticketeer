package git

import (
	"errors"
	"strings"
)

const headPrefix = "ref: refs/heads/"

// ErrBranchNotFound is returned when repository HEAD is in detached state
var ErrBranchNotFound = errors.New("HEAD is not on a branch")

// BranchNameFromHead extracts branch name from raw HEAD content
func BranchNameFromHead(head string) (string, error) {
	if !strings.HasPrefix(head, headPrefix) {
		return "", ErrBranchNotFound
	}
	branch := strings.TrimPrefix(head, headPrefix)
	branch = strings.TrimRight(branch, " \n")

	return branch, nil
}
