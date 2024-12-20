package hook_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/mishamyrt/ticketeer/internal/hook"
	"github.com/mishamyrt/ticketeer/internal/testutil"
)

func TestDetectRunner(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		path    string
		want    *hook.Runner
		wantErr bool
	}{
		{"", nil, false},
		{"../../testdata/hook_runners/lefthook.sh", &hook.LefthookRunner, false},
		{"../../testdata/hook_runners/ticketeer.sh", &hook.TicketeerRunner, false},
		{"../../testdata/hook_runners/unknown.sh", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got, err := hook.DetectRunner(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectRunner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DetectRunner() got = %v, want %v", got, tt.want)
			}
		})
	}

	// Create 2 MB file
	dir := t.TempDir()
	binaryPath := filepath.Join(dir, "binary")
	payload := testutil.RandomBytes(1024 * 1024 * 2)
	err := os.WriteFile(binaryPath, payload, 0644)
	if err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	_, err = hook.DetectRunner(binaryPath)
	if !errors.Is(err, hook.ErrUnknownRunner) {
		t.Errorf("DetectRunner() got = %v, want %v", err, hook.ErrUnknownRunner)
	}
}
