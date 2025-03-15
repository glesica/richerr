package richerr

import (
	"errors"
	"fmt"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestAllFields(t *testing.T) {
	t.Run("should handle a nil error", func(t *testing.T) {
		fields := AllFields(nil)
		assert.Equal(t, nil, fields)
	})

	t.Run("should handle a wrapped error", func(t *testing.T) {
		richErr := New("richerr").WithField("foo", "bar")
		wrapErr := fmt.Errorf("wrapper: %w", richErr)
		fields := AllFields(wrapErr)
		assert.Equal(t, Fields{{"foo", "bar"}}, fields)
	})

	t.Run("should handle multiple wrapped errors", func(t *testing.T) {
		richErr0 := New("richerr0").WithField("foo", "bar")
		richErr1 := Wrap(richErr0, "richerr1").WithField("baz", "fuz")
		wrapErr := fmt.Errorf("wrapper: %w", richErr1)
		fields := AllFields(wrapErr)
		assert.Equal(t, Fields{{"baz", "fuz"}, {"foo", "bar"}}, fields)
	})

	t.Run("should handle joined errors", func(t *testing.T) {
		richErr0 := New("richerr0").WithField("foo", "bar")
		richErr1 := New("richerr1").WithField("baz", "fuz")
		joinErr := errors.Join(richErr0, richErr1)
		fields := AllFields(joinErr)
		assert.Equal(t, Fields{{"foo", "bar"}, {"baz", "fuz"}}, fields)
	})
}
