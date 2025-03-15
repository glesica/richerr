package richerr

import (
	"slices"
)

// Error is an error implementation that supports adding
// arbitrary metadata.
type Error struct {
	error

	fields Fields
}

func (e Error) Unwrap() error {
	return e.error
}

func (e Error) Fields() Fields {
	if e.fields == nil {
		return nil
	}

	return slices.Clone(e.fields)
}

// WithField adds a single field with the given name and
// value and returns the updated Error.
func (e Error) WithField(name string, value any) Error {
	e.fields = append(e.fields, Field{Name: name, Value: value})

	return e
}

// WithFields adds multiple fields and returns the updated
// Error.
func (e Error) WithFields(fields Fields) Error {
	e.fields = append(e.fields, fields...)

	return e
}
