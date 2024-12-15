package hook

import (
	"os"
)

// ScriptType is a hook script type
type ScriptType interface {
	Name() string
	UsageGuide() string
	assert(content string) bool
}

var knownTypes = []ScriptType{
	ticketeerScript{},
	lefthookScript{},
}

// FindType returns hook script type
func FindType(hookPath string) ScriptType {
	content, err := os.ReadFile(hookPath)
	if err != nil {
		return nil
	}
	for _, scriptType := range knownTypes {
		if scriptType.assert(string(content)) {
			return scriptType
		}
	}
	return nil
}
