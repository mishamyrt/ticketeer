package hook

import "strings"

type lefthookScript struct{}

var _ ScriptType = lefthookScript{}

func (l lefthookScript) Name() string {
	return "lefthook"
}

func (l lefthookScript) UsageGuide() string {
	return "lefthook"
}

func (l lefthookScript) assert(content string) bool {
	return strings.Contains(content, "LEFTHOOK_BIN")
}
