package hook

import "path/filepath"

// Name is git hook file name
const Name = "prepare-commit-msg"

// Path returns path to git hook
func Path(hooksDir string) string {
	return filepath.Join(hooksDir, Name)
}
