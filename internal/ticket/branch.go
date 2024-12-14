package ticket

import (
	"errors"
	"regexp"
)

// ErrInvalidBranchName is returned when branch name is invalid
var ErrInvalidBranchName = errors.New("invalid branch name")

var branchFormatRe = regexp.MustCompile("^.[^/]*/(.[^/]*).*$")

// ParseFromBranchName parses ticket id from branch name
func ParseFromBranchName(branchName string, format Format) (ID, error) {
	match := branchFormatRe.FindStringSubmatch(branchName)
	if len(match) < 2 || len(match[1]) < 2 {
		return "", ErrInvalidBranchName
	}
	return ParseID(match[1], format)
}
