package service

import (
	"errors"

	"github.com/djsilvajr/go-skeleton/internal/domain/user/service/usecase"
)

// Sentinel errors — re-exported for type-safe checks in tests and handlers.
var (
	ErrInvalidCredentials = usecase.ErrInvalidCredentials
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyInUse  = errors.New("email already in use")
)
