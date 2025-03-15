package richerr

import (
	"errors"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should create a new error", func(t *testing.T) {
		err := New("test error")
		assert.Equal(t, "test error", err.Error())
	})
}

func TestWrap(t *testing.T) {
	t.Run("should create a wrapped error", func(t *testing.T) {
		err := Wrap(errors.New("test error"), "wrapper")
		assert.Equal(t, "wrapper: test error", err.Error())
	})
}
