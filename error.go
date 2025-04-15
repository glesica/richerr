package richerr

import (
	"slices"
)

// Error is an error implementation that supports adding
// arbitrary metadata.
type Error struct {
	error

	scope  string
	fields Fields
}

func (e Error) Unwrap() error {
	return e.error
}

func (e Error) Scope() string {
	return e.scope
}

// Fields returns the fields associated with this error,
// specifically. It ignores any nested errors.
func (e Error) Fields() Fields {
	return slices.Clone(e.fields)
}

// WithScope sets the scope of this error and returns the
// updated error.
func (e Error) WithScope(scope string) Error {
	e.scope = scope

	return e
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
