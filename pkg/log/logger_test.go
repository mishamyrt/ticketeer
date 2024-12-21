package log_test

import (
	"io"
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

type stdLoggerWrapper struct{}

var _ log.Logger = stdLoggerWrapper{}

func (stdLoggerWrapper) SetOutput(w io.Writer) {
	log.SetOutput(w)
}
func (stdLoggerWrapper) SetLevel(level log.Level) {
	log.SetLevel(level)
}
func (stdLoggerWrapper) Debug(v ...any) {
	log.Debug(v...)
}
func (stdLoggerWrapper) Debugf(format string, a ...any) {
	log.Debugf(format, a...)
}
func (stdLoggerWrapper) Info(v ...any) {
	log.Info(v...)
}
func (stdLoggerWrapper) Infof(format string, a ...any) {
	log.Infof(format, a...)
}
func (stdLoggerWrapper) Error(v ...any) {
	log.Error(v...)
}
func (stdLoggerWrapper) Errorf(format string, a ...any) {
	log.Errorf(format, a...)
}

func TestPureLogger(t *testing.T) {
	t.Parallel()
	Suite(t, log.NewPure())
	Suite(t, &stdLoggerWrapper{})
}
