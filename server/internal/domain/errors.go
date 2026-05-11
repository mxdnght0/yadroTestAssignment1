package domain

import "errors"

var (
	ErrFileIsNotCreated = errors.New("file is not created")
	ErrDNSAlreadyExists = errors.New("dns already exists")
	ErrDNSNotFound      = errors.New("dns not found")
)
