package richerr

// A Field is a single name-value pair that can be added to
// an Error. If a scope is assigned, it will be used to indicate
// the nesting of fields from wrapped errors.
type Field struct {
	Name  string `json:"name,required"`
	Value any    `json:"value,required"`
}

// Fields is a slice of name-value pairs.
type Fields []Field
