package hook

import "path/filepath"

// Path returns path to git hook
func Path(name string, hooksDir string) string {
	return filepath.Join(hooksDir, name)
}
