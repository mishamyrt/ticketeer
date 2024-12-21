package log

import "io"

// Level represents log level
type Level int

const (
	// LevelError represents error log level
	LevelError Level = iota

	// LevelInfo represents info log level
	LevelInfo

	// LevelDebug represents debug log level
	LevelDebug
)

// Int returns level as int
func (l Level) Int() int {
	return int(l)
}

// Logger represents app logger interface
type Logger interface {
	SetOutput(w io.Writer)
	SetLevel(level Level)

	Debug(v ...any)
	Debugf(format string, a ...any)
	Info(v ...any)
	Infof(format string, a ...any)
	Error(v ...any)
	Errorf(format string, a ...any)
}
