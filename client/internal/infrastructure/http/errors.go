package http

import "errors"

var (
	ErrInvalidIP      = errors.New("dns address is not a valid IP")
	ErrAlreadyExists  = errors.New("dns already exists")
	ErrNotFound       = errors.New("dns not found")
	ErrFileNotCreated = errors.New("dns file not found on server")
)
