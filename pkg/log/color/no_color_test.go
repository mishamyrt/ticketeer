package color_test

import (
	"testing"

	"github.com/mishamyrt/ticketeer/pkg/log/color"
)

func TestSetNoColor(t *testing.T) {
	color.SetNoColor(false)

	colored := color.Red("test")

	color.SetNoColor(true)
	noColor := color.Red("test")

	if len(colored) <= len(noColor) {
		t.Errorf("Expected %q to be longer than %q", colored, noColor)
	}
}

func TestWrapNoColor(t *testing.T) {
	isCalled := false

	callback := func(a ...any) string {
		isCalled = true
		return ""
	}
	wrappedColorize := color.WrapNoColor(callback)

	color.SetNoColor(true)
	_ = wrappedColorize("test")
	if isCalled {
		t.Error("Expected callback not to be called")
	}

	color.SetNoColor(false)
	_ = wrappedColorize("test")
	if !isCalled {
		t.Error("Expected callback to be called")
	}
}
