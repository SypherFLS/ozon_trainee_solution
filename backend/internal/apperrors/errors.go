package apperrors

import "errors"

var (
    ErrNotFound   = errors.New("not found")
    ErrInvalidURL = errors.New("invalid url")
	ErrConflict = errors.New("conflict")
    ErrTooManyAttempts = errors.New("too many attempts")
)