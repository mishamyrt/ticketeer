package log

import (
	"io"
)

var std = New()

// SetLevel sets global log level
func SetLevel(level Level) {
	std.SetLevel(level)
}

// SetOut sets global log output
func SetOut(out io.Writer) {
	std.SetOut(out)
}

// Error writes error message
func Error(message string) {
	std.Error(message)
}

// Errorf writes formatted error message
func Errorf(format string, a ...any) {
	std.Errorf(format, a...)
}

// Warn writes warning message
func Warn(message string) {
	std.Warn(message)
}

// Warnf writes formatted warning message
func Warnf(format string, a ...any) {
	std.Warnf(format, a...)
}

// Info writes info message
func Info(message string) {
	std.Info(message)
}

// Infof writes formatted info message
func Infof(format string, a ...any) {
	std.Infof(format, a...)
}

// Debug writes debug message
func Debug(message string) {
	std.Debug(message)
}

// Debugf writes formatted debug message
func Debugf(format string, a ...any) {
	std.Debugf(format, a...)
}

// Print writes info message
func Print(message string) {
	std.Print(message)
}

// Printf writes formatted info message
func Printf(format string, a ...any) {
	std.Printf(format, a...)
}

// Success writes success message
func Success(message string) {
	std.Success(message)
}
