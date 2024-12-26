package ticketeer

import (
	"errors"
	"fmt"
)

var errHandled = errors.New("[handled]")

// IsHandledError checks if error is already handled by application
func IsHandledError(err error) bool {
	return errors.Is(err, errHandled)
}

func wrapHandledError(err error) error {
	return fmt.Errorf("%w %s", errHandled, err)
}
