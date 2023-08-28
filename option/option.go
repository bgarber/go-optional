// (C) Copyright 2023 Bryan Garber

package option

import "errors"

// List of errors
var ErrFailedUnwrap = errors.New("failed unwrap")

// Optional is a container struct, which can assume the value of None or Some(x)
type Optional[T any] struct {
	content any
}

// Some is a function for creating an Optional that is something
func Some[T any](v T) Optional[T] {
	return Optional[T]{
		content: v,
	}
}

// None checks if an Optional is nothing
func (o Optional[T]) None() bool {
	return o.content == nil
}

// Unwrap unwraps the Optional content to the expected type
func (o Optional[T]) Unwrap() (T, error) {
	var v T
	v, ok := o.content.(T)
	if !ok {
		return v, ErrFailedUnwrap
	}

	return v, nil
}

// UnwrapOr tries to unwrap the Optional. If it's None, then returns the
// provided value.
func (o Optional[T]) UnwrapOr(def T) T {
	if o.content == nil {
		return def
	}

	// it is safe to discard any error
	v, _ := o.Unwrap()
	return v
}
