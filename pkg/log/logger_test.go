package log_test

import (
	"strings"
	"testing"

	"github.com/mishamyrt/ticketeer/pkg/log"
)

type PrintFunc func(v ...any)
type PrintfFunc func(format string, v ...any)

func AssertPrintFunc(t *testing.T, printFunc PrintFunc, buf *strings.Builder) bool {
	buf.Reset()
	printFunc("test")
	if !strings.Contains(buf.String(), "test") {
		t.Errorf("Expected %q to contain %q", buf.String(), "test")
		return false
	}

	return true
}

func AssertPrintfFunc(t *testing.T, printfFunc PrintfFunc, buf *strings.Builder) bool {
	buf.Reset()
	printfFunc("^%s$", "test")
	if !strings.Contains(buf.String(), "^test$") {
		t.Errorf("Expected %q to contain %q", buf.String(), "^test$")
		return false
	}

	return true
}

func AssertLevel(t *testing.T, logger log.Logger, buf *strings.Builder) bool {
	buf.Reset()
	logger.SetLevel(log.LevelInfo)
	logger.Debug("test")
	if buf.Len() != 0 {
		t.Error("Expected buffer to be empty")
		return false
	}
	logger.SetLevel(log.LevelDebug)
	logger.Debug("test")
	if buf.Len() == 0 {
		t.Error("Expected buffer to not be empty")
		return false
	}

	return true
}

func AssertOutputSet(t *testing.T, logger log.Logger, buf *strings.Builder) bool {
	buf.Reset()
	logger.SetOutput(buf)
	logger.Info("test")
	if buf.Len() == 0 {
		t.Error("Expected buffer to not be empty")
		return false
	}

	return true
}

func Suite(t *testing.T, logger log.Logger) {
	var buf strings.Builder

	if !AssertOutputSet(t, logger, &buf) {
		return
	}

	if !AssertLevel(t, logger, &buf) {
		return
	}

	logger.SetLevel(log.LevelDebug)

	printFuncs := []PrintFunc{
		logger.Debug,
		logger.Info,
		logger.Error,
	}
	for _, printFunc := range printFuncs {
		AssertPrintFunc(t, printFunc, &buf)
	}

	printfFuncs := []PrintfFunc{
		logger.Debugf,
		logger.Infof,
		logger.Errorf,
	}
	for _, printfFunc := range printfFuncs {
		AssertPrintfFunc(t, printfFunc, &buf)
	}
}

func TestPureLogger(t *testing.T) {
	t.Parallel()
	Suite(t, log.NewPure())
}
