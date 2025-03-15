package richerr

import (
	"errors"
	"github.com/alecthomas/assert/v2"
	"slices"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("should allow access to the underlying error", func(t *testing.T) {
		err := Error{error: errors.New("test error")}
		assert.Equal(t, "test error", err.Error())
	})
}

func TestError_Fields(t *testing.T) {
	t.Run("should provide access to the fields", func(t *testing.T) {
		f := Fields{
			{"foo", "bar"},
			{"baz", "qux"},
		}
		err := Error{
			error:  errors.New("test error"),
			fields: slices.Clone(f),
		}
		assert.Equal(t, f, err.Fields())
	})
}

func TestError_Unwrap(t *testing.T) {
	t.Run("should unwrap to the underlying error", func(t *testing.T) {
		origErr := errors.New("test error")
		err := Error{error: origErr}
		assert.Equal(t, origErr, err.Unwrap())
	})
}

func TestError_AddField(t *testing.T) {
	t.Run("should add a new field", func(t *testing.T) {
		err := Error{error: errors.New("test error")}
		err = err.WithField("foo", "bar")
		assert.Equal(t, Fields{{"foo", "bar"}}, err.fields)
	})
}
