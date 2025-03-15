package richerr

import (
	"errors"
	"fmt"
	"strings"
)

// New creates a new error with the given message and applies
// the given options.
func New(msg string) Error {
	richErr := Error{
		error: errors.New(msg),
	}

	return richErr
}

// Wrap creates a new rich error that wraps the given error,
// with the given options applied.
func Wrap(err error, msg string) Error {
	if !strings.HasSuffix(msg, ":") {
		msg += ":"
	}

	richErr := Error{
		error: fmt.Errorf("%s %w", msg, err),
	}

	return richErr
}
