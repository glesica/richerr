package richerr

// Collect extracts the union of the fields found on all
// rich errors within the error tree with the given root. Traversal
// is depth-first, which affects the order of the fields in the
// resulting slice but otherwise has no effect. Multiple fields with
// the same name are preserved. Field names will be namespaced with
// the text of the error with which they are associated, or the
// scopes for each error, if provided.
func Collect(err error) Fields {
	return collect(err, "")
}

func collect(err error, scope Scope) Fields {
	if err == nil {
		return nil
	}

	// If we've found an Error, or anything implementing an
	// equivalent Fields() method, then grab its fields before
	// traversing the rest of the tree.
	var fields Fields
	if fieldsErr, fieldsOK := err.(interface{ Fields() Fields }); fieldsOK {
		fields = fieldsErr.Fields()

		// If this level of nesting introduced new fields then,
		// and only then, it will also introduce a new scope.
		// This way we don't add empty scopes for error wrappers,
		// like those from fmt.Errorf.
		var nextScopeLevel Scope
		if scopeErr, scopeOK := err.(interface{ Scope() Scope }); scopeOK {
			nextScopeLevel = scopeErr.Scope()
		} else {
			nextScopeLevel = Scope(err.Error())
		}

		if scope == "" {
			scope = nextScopeLevel
		} else {
			scope += "/" + nextScopeLevel
		}

		scopeFields(fields, scope)
	}

	// The traversal strategy here is cribbed from errors.As,
	// adapted to recover fields from the entire tree.
	switch wrapErr := err.(type) {
	case interface{ Unwrap() error }:
		err := wrapErr.Unwrap()
		if err != nil {
			fields = append(fields, collect(err, scope)...)
		}
	case interface{ Unwrap() []error }:
		for _, err := range wrapErr.Unwrap() {
			if err != nil {
				fields = append(fields, collect(err, scope)...)
			}
		}
	}

	return fields
}

func scopeFields(fields Fields, scope Scope) {
	if scope == "" {
		return
	}

	for index, field := range fields {
		field.Name = string(scope) + "/" + field.Name
		fields[index] = field
	}
}
