package log

import "io"

var std = New()

// SetOutput sets output
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// SetLevel sets log level
func SetLevel(level Level) {
	std.SetLevel(level)
}

// Debug writes debug message
func Debug(v ...any) {
	std.Debug(v...)
}

// Debugf writes formatted debug message
func Debugf(format string, a ...any) {
	std.Debugf(format, a...)
}

// Info writes info message
func Info(v ...any) {
	std.Info(v...)
}

// Infof writes formatted info message
func Infof(format string, a ...any) {
	std.Infof(format, a...)
}

// Error writes error message
func Error(v ...any) {
	std.Error(v...)
}

// Errorf writes formatted error message
func Errorf(format string, a ...any) {
	std.Errorf(format, a...)
}
