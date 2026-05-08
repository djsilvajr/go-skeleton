package errors

import "errors"

var (
	// ErrInvalidCredentials is returned when email/password do not match.
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrUserNotFound is returned when a user lookup yields no rows.
	ErrUserNotFound = errors.New("user not found")

	// ErrEmailAlreadyInUse is returned when creating a user with an email
	// that already exists in the database.
	ErrEmailAlreadyInUse = errors.New("email already in use")

	// ErrNameAlreadyInUse is returned when creating a user with a name
	// that already exists in the database.
	ErrNameAlreadyInUse = errors.New("name already in use")
)
