// (C) Copyright 2023 Bryan Garber

package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptional(t *testing.T) {
	t.Run("test Optional with something", func(t *testing.T) {
		theInt := 42
		someInt := Some(theInt)

		unwrapped, err := someInt.Unwrap()
		assert.NoError(t, err)
		assert.Equal(t, theInt, unwrapped)
	})

	t.Run("test Optional with none", func(t *testing.T) {
		var nothing Optional[int]

		assert.Equal(t, nothing.None(), true)
	})

	t.Run("test invalid unwrap", func(t *testing.T) {
		var nothing Optional[int]

		_, err := nothing.Unwrap()
		assert.ErrorIs(t, err, ErrFailedUnwrap)
	})

	t.Run("test unwrap with default value", func(t *testing.T) {
		var (
			defaultInt int = 42
			otherInt   int = 7
			nothing    Optional[int]
			someInt    Optional[int] = Some(otherInt)
		)

		value := nothing.UnwrapOr(defaultInt)
		assert.Equal(t, defaultInt, value)

		value = someInt.UnwrapOr(defaultInt)
		assert.Equal(t, otherInt, value)
	})
}
