package richerr

import (
	"errors"
	"fmt"
	"github.com/alecthomas/assert/v2"
	"testing"
)

func TestCollect(t *testing.T) {
	t.Run("should handle a nil error", func(t *testing.T) {
		fields := Collect(nil)
		assert.Equal(t, nil, fields)
	})

	t.Run("should handle a wrapped error", func(t *testing.T) {
		richErr := New("richerr").WithField("foo", "bar")
		wrapErr := fmt.Errorf("wrapper: %w", richErr)
		fields := Collect(wrapErr)
		assert.Equal(t, Fields{
			{"foo", "bar"},
		}, fields)
	})

	t.Run("should handle multiple wrapped errors", func(t *testing.T) {
		richErr0 := New("richerr0").WithField("foo", "bar")
		richErr1 := Wrap(richErr0, "richerr1").WithField("baz", "fuz")
		wrapErr := fmt.Errorf("wrapper: %w", richErr1)
		fields := Collect(wrapErr)
		assert.Equal(t, Fields{
			{"baz", "fuz"},
			{"foo", "bar"},
		}, fields)
	})

	t.Run("should handle joined errors", func(t *testing.T) {
		richErr0 := New("richerr0").WithField("foo", "bar")
		richErr1 := New("richerr1").WithField("baz", "fuz")
		joinErr := errors.Join(richErr0, richErr1)
		fields := Collect(joinErr)
		assert.Equal(t, Fields{
			{"foo", "bar"},
			{"baz", "fuz"},
		}, fields)
	})

	t.Run("should handle scope on a single error", func(t *testing.T) {
		richErr := New("richerr").
			WithScope("rich_scope").
			WithField("foo", "bar")
		fields := Collect(richErr)
		assert.Equal(t, Fields{
			{"rich_scope/foo", "bar"},
		}, fields)
	})

	t.Run("should handle scope on a single wrapped error", func(t *testing.T) {
		richErr := New("richerr").
			WithScope("rich_scope").
			WithField("foo", "bar")
		wrapErr := fmt.Errorf("wrapper: %w", richErr)
		fields := Collect(wrapErr)
		assert.Equal(t, Fields{
			{"rich_scope/foo", "bar"},
		}, fields)
	})

	t.Run("should handle scope on multiple errors", func(t *testing.T) {
		richErr0 := New("richerr0").
			WithScope("rich_scope0").
			WithField("foo", "bar")
		richErr1 := Wrap(richErr0, "richerr1").
			WithScope("rich_scope1").
			WithField("baz", "fuz")
		wrapErr := fmt.Errorf("wrapper: %w", richErr1)
		fields := Collect(wrapErr)
		assert.Equal(t, Fields{
			{"rich_scope1/baz", "fuz"},
			{"rich_scope1/rich_scope0/foo", "bar"},
		}, fields)
	})

	t.Run("should handle scope on joined errors", func(t *testing.T) {
		richErr0 := New("richerr0").
			WithScope("rich_scope0").
			WithField("foo", "bar")
		richErr1 := New("richerr1").
			WithScope("rich_scope1").
			WithField("baz", "fuz")
		joinErr := errors.Join(richErr0, richErr1)
		fields := Collect(joinErr)
		assert.Equal(t, Fields{
			{"rich_scope0/foo", "bar"},
			{"rich_scope1/baz", "fuz"},
		}, fields)
	})
}
