package service

import "errors"

// Sentinel errors — use these for type-safe error checks in tests and handlers.
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyInUse  = errors.New("email already in use")
)
