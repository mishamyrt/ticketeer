package log

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/fatih/color"
)

// Level represents log level
type Level int

//revive:disable
const (
	LevelError Level = iota
	LevelInfo
	LevelWarn
	LevelDebug
)

//revive:enable

// Logger represents app logger
type Logger struct {
	colored bool
	level   Level
	out     io.Writer
	lock    sync.Mutex
}

// New creates new logger
func New() *Logger {
	return &Logger{
		level:   LevelWarn,
		out:     os.Stdout,
		colored: true,
	}
}

// SetLevel sets log level
func (l *Logger) SetLevel(level Level) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.level = level
}

// SetColored sets colored output
func (l *Logger) SetColored(colored bool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.colored = colored
}

// SetOut sets output
func (l *Logger) SetOut(out io.Writer) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.out = out
}

// Error writes error message
func (l *Logger) Error(message string) {
	l.Log(LevelError, message, color.FgRed)
}

// Errorf writes formatted error message
func (l *Logger) Errorf(format string, a ...any) {
	l.Error(fmt.Sprintf(format, a...))
}

// Warn writes warning message
func (l *Logger) Warn(message string) {
	l.Log(LevelWarn, message, color.FgYellow)
}

// Warnf writes formatted warning message
func (l *Logger) Warnf(format string, a ...any) {
	l.Warn(fmt.Sprintf(format, a...))
}

// Info writes info message
func (l *Logger) Info(message string) {
	l.Log(LevelInfo, message, color.FgCyan)
}

// Infof writes formatted info message
func (l *Logger) Infof(format string, a ...any) {
	l.Info(fmt.Sprintf(format, a...))
}

// Debug writes debug message
func (l *Logger) Debug(message string) {
	l.Log(LevelDebug, message, color.Faint)
}

// Debugf writes formatted debug message
func (l *Logger) Debugf(format string, a ...any) {
	l.Debug(fmt.Sprintf(format, a...))
}

// Print writes info message
func (l *Logger) Print(message string) {
	l.Log(LevelInfo, message)
}

// Printf writes formatted info message
func (l *Logger) Printf(format string, a ...any) {
	l.Print(fmt.Sprintf(format, a...))
}

// Success writes success message
func (l *Logger) Success(message string) {
	l.Log(LevelInfo, message, color.FgGreen)
}

// Log writes message
func (l *Logger) Log(level Level, message string, attributes ...color.Attribute) {
	if level > l.level {
		return
	}
	color.
		New(attributes...).
		SetWriter(l.out).
		Println(message)
}
