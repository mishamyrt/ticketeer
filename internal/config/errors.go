package config

import (
	"errors"
	"fmt"
)

var (
	ErrFileNotFound    = errors.New("config file not found")
	ErrUnknownLocation = errors.New("unknown ticket location")
)

func newErrFileNotFound(path string) error {
	return fmt.Errorf("%w at %s", ErrFileNotFound, path)
}
