package hook

import (
	"errors"
	"os"
	"strings"
)

const runnerSizeLimit = 1024 * 1024 // 1MB

var (
	// ErrUnknownRunner is returned when hook runner is unknown
	ErrUnknownRunner = errors.New("unable to detect git hook runner from file content")
)

// Runner represents git hook runner
type Runner struct {
	// Name of the hook runner.
	Name string
	// GuideAnchor is the anchor for the usage guide at README.
	GuideAnchor string

	// assert is a function that returns true if the hook runner is detected.
	assert func(content string) bool
}

var (
	// LefthookRunner represents lefthook runner
	LefthookRunner = Runner{
		Name:        "lefthook",
		GuideAnchor: "lefthook",
		// Detects lefthook runner by checking for its binary environment variable
		assert: func(content string) bool {
			return strings.Contains(content, "LEFTHOOK_BIN")
		},
	}

	// TicketeerRunner represents ticketeer runner
	TicketeerRunner = Runner{
		Name:        "ticketeer",
		GuideAnchor: "",
		// Detects ticketeer runner by checking for its binary environment variable
		assert: func(content string) bool {
			return strings.Contains(content, "TICKETEER_BIN")
		},
	}
)

var runners = []*Runner{
	&TicketeerRunner,
	&LefthookRunner,
}

// DetectRunner returns hook runner based on hook content
func DetectRunner(path string) (*Runner, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	if info.Size() > runnerSizeLimit { // 1MB limit
		return nil, ErrUnknownRunner
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	for _, runner := range runners {
		if runner.assert(string(content)) {
			return runner, nil
		}
	}
	return nil, ErrUnknownRunner
}
