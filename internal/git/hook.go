package git

import (
	"strings"
)

// HooksPath returns path to git hooks directory
func (r *Repository) HooksPath() (string, error) {
	err := AssertRepository(r.path)
	if err != nil {
		return "", err
	}
	out, err := Exec(r.path, "rev-parse", "--git-path", "hooks")
	if err != nil {
		return "", err
	}
	return strings.Trim(out, "\n"), nil
}
