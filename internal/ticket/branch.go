package ticket

import (
	"errors"
	"regexp"
)

// ErrInvalidBranchName is returned when branch name is invalid
var ErrInvalidBranchName = errors.New("invalid branch name")

// BranchFormat represents git branch name format
type BranchFormat string

var (
	// GitFlowBranch represents git-flow branch format e.g. feature/FEAT-123[/description]
	GitFlowBranch BranchFormat = "^.[^/]*/(.[^\n/]*).*$"
	// GitFlowTypelessBranch represents git-flow branch format without type e.g. FEAT-123[/description]
	GitFlowTypelessBranch BranchFormat = "(.[^\n/]*).*$"
	// TicketIDBranch represents ticket id branch format e.g. FEAT-123
	TicketIDBranch BranchFormat = "^([A-Za-z0-9-#]*)$"
)

// FindInBranch parses ticket id from branch name
func FindInBranch(branchName string, format BranchFormat) (string, error) {
	re, err := regexp.Compile(string(format))
	if err != nil {
		return "", err
	}
	match := re.FindStringSubmatch(branchName)
	if len(match) < 2 || len(match[1]) < 2 {
		return "", ErrInvalidBranchName
	}
	return match[1], nil
}
