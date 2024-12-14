package ticket

import (
	"errors"
	"regexp"
)

var branchFormatRe = regexp.MustCompile("^.[^/]*/(.[^/]*).*$")

var ErrInvalidBranchName = errors.New("invalid branch name")

func ParseFromBranchName(branchName string, format Format) (ID, error) {
	match := branchFormatRe.FindStringSubmatch(branchName)
	if len(match) < 2 || len(match[1]) < 2 {
		return "", ErrInvalidBranchName
	}
	return ParseID(match[1], format)
}
