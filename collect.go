package richerr

// Collect extracts the union of the fields found on all
// rich errors within the error tree with the given root. Traversal
// is depth-first, which affects the order of the fields in the
// resulting slice but otherwise has no effect. Multiple fields with
// the same name are preserved. Field names will be namespaced with
// the text of the error with which they are associated, or the
// scopes for each error, if provided.
func Collect(err error) []Fields {
	if err == nil {
		return nil
	}

	var fields []Fields

	// If we've found an Error, or anything implementing an
	// equivalent Fields() method, then grab its fields before
	// traversing the rest of the tree.
	if fieldsErr, fieldsOK := err.(interface{ Fields() Fields }); fieldsOK {
		theseFields := fieldsErr.Fields()
		if theseFields != nil {
			fields = append(fields, theseFields)
		}
	}

	// The traversal strategy here is cribbed from errors.As,
	// adapted to recover fields from the entire tree.
	switch wrapErr := err.(type) {
	case interface{ Unwrap() error }:
		wrappedErr := wrapErr.Unwrap()
		if wrappedErr != nil {
			fields = append(fields, Collect(wrappedErr)...)
		}
	case interface{ Unwrap() []error }:
		for _, wrappedErr := range wrapErr.Unwrap() {
			if wrappedErr != nil {
				fields = append(fields, Collect(wrappedErr)...)
			}
		}
	}

	return fields
}
