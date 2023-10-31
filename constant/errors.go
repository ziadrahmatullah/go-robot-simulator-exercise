package constant

import "errors"

var(
	ErrInvalidPosition = errors.New("invalid position")
	ErrInvalidMovement = errors.New("invalid movement")
)