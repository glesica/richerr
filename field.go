package richerr

import (
	"fmt"
)

// A Field is a single name-value pair that can be added to
// an Error.
type Field struct {
	Name  string
	Value any
}

func (f Field) String() string {
	return fmt.Sprintf("%s: %v", f.Name, f.Value)
}

type Fields []Field

// AllFields extracts the union of the fields found on all
// rich errors within the error tree with the given root. Traversal
// is depth-first, which affects the order of the fields in the
// resulting slice but otherwise has no effect.
func AllFields(err error) Fields {
	if err == nil {
		return nil
	}

	// If we've found an Error, or anything implementing an
	// equivalent Fields() method, then grab its fields before
	// traversing the rest of the tree.
	var fields Fields
	if fieldsErr, ok := err.(interface{ Fields() Fields }); ok {
		fields = fieldsErr.Fields()
	}

	// The traversal strategy here is cribbed from errors.As,
	// adapted to recover fields from the entire tree.
	switch wrapErr := err.(type) {
	case interface{ Unwrap() error }:
		err = wrapErr.Unwrap()
		if err != nil {
			fields = append(fields, AllFields(err)...)
		}
	case interface{ Unwrap() []error }:
		for _, err := range wrapErr.Unwrap() {
			if err != nil {
				fields = append(fields, AllFields(err)...)
			}
		}
	}

	return fields
}
