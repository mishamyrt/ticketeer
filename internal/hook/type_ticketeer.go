package hook

import "strings"

type ticketeerScript struct{}

var _ ScriptType = ticketeerScript{}

func (l ticketeerScript) Name() string {
	return "ticketeer"
}

func (l ticketeerScript) UsageGuide() string {
	return "ticketeer"
}

func (l ticketeerScript) assert(content string) bool {
	return strings.Contains(content, "TICKETEER_BIN")
}
