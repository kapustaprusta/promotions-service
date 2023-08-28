package repository

import "errors"

var (
	// ErrNilValue encapsulates nil value error
	ErrNilValue = errors.New("nil value")

	// ErrRecordNotFound encapsulates record not found error
	ErrRecordNotFound = errors.New("record not found")
)
