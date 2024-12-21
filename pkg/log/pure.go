package log

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/mishamyrt/ticketeer/pkg/log/color"
)

// PureLogger represents app logger
type PureLogger struct {
	level Level
	out   io.Writer
	mu    sync.Mutex
}

// NewPure creates new pure logger
func NewPure() Logger {
	return &PureLogger{
		level: LevelInfo,
		out:   os.Stdout,
	}
}

// New creates new default logger
func New() Logger {
	return NewPure()
}

// SetOutput sets output
func (l *PureLogger) SetOutput(out io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.out = out
}

// SetLevel sets log level
func (l *PureLogger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.level = level
}

// Error writes error message
func (l *PureLogger) Error(v ...any) {
	l.log(LevelError, color.Red(v...))
}

// Errorf writes formatted error message
func (l *PureLogger) Errorf(format string, a ...any) {
	l.logf(LevelError, color.Red(format), a...)
}

// Info writes info message
func (l *PureLogger) Info(v ...any) {
	l.log(LevelInfo, v...)
}

// Infof writes formatted info message
func (l *PureLogger) Infof(format string, a ...any) {
	l.logf(LevelInfo, format, a...)
}

// Debug writes debug message
func (l *PureLogger) Debug(v ...any) {
	l.log(LevelDebug, color.Dim(v...))
}

// Debugf writes formatted debug message
func (l *PureLogger) Debugf(format string, a ...any) {
	l.logf(LevelDebug, color.Dim(format), a...)
}

// Print writes info message
func (l *PureLogger) Print(v ...any) {
	l.log(LevelInfo, v...)
}

// Printf writes formatted info message
func (l *PureLogger) Printf(format string, a ...any) {
	l.logf(LevelInfo, format, a...)
}

func (l *PureLogger) log(level Level, v ...any) {
	l.logf(level, "%s", fmt.Sprint(v...))
}

func (l *PureLogger) logf(level Level, format string, a ...any) {
	if level.Int() > l.level.Int() {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	fmt.Fprintf(l.out, format+"\n", a...)
}
