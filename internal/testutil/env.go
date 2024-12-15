package testutil

import "os"

// EnvMock mocks environment variable
type EnvMock struct {
	variable      string
	previousValue string
	isApplied     bool
}

// IsApplied returns true if environment variable is mocked
func (e *EnvMock) IsApplied() bool {
	return e.isApplied
}

// Restore environment variable to original value
func (e *EnvMock) Restore() {
	os.Setenv(e.variable, e.previousValue)
	e.isApplied = false
}

// Set environment variable
func (e *EnvMock) Set(value string) {
	if !e.isApplied {
		e.previousValue = os.Getenv(e.variable)
	}
	e.isApplied = true
	os.Setenv(e.variable, value)
}

// NewEnvMock creates new environment variable mock
func NewEnvMock(variable string, value string) *EnvMock {
	mock := &EnvMock{
		variable: variable,
	}
	mock.Set(value)
	return mock
}
