package domain

import "errors"

var (
	// ErrNilValue encapsulates nil value error
	ErrNilValue = errors.New("nil value")

	// ErrInvalidValue encapsulates invalid value error
	ErrInvalidValue = errors.New("invalid value")
)
