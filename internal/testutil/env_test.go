package testutil_test

import (
	"os"
	"testing"

	"github.com/mishamyrt/ticketeer/internal/testutil"
)

func TestEnvMock(t *testing.T) {
	t.Parallel()

	const testVar = "FOO"
	const testValue = "BAR"

	os.Setenv(testVar, testValue)
	if os.Getenv(testVar) != testValue {
		t.Fatalf("expected %s to be %s, got %s", testVar, testValue, os.Getenv(testVar))
	}
	mock := testutil.NewEnvMock(testVar, "")
	if os.Getenv(testVar) != "" {
		t.Fatalf("expected %s to be empty, got %s", testVar, os.Getenv(testVar))
	}
	if !mock.IsApplied() {
		t.Fatal("expected mock to be applied")
	}
	mock.Restore()
	if os.Getenv(testVar) != testValue {
		t.Fatalf("expected %s to be %s, got %s", testVar, testValue, os.Getenv(testVar))
	}
	if mock.IsApplied() {
		t.Fatal("expected mock to not be applied")
	}
}
