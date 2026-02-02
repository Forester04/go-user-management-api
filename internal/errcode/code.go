package errcode

import (
	"errors"
	"fmt"
)

// GoCleanError used in this project
type GoCleanError struct {
	error
	int
}

var (
	// generic errors (100-199)
	ErrUndefined      = newErrcode("undefined error", 100)
	ErrNotImplemented = newErrcode("not implemented", 101)

	// database errors (200-299)
	ErrDatabase          = newErrcode("database error", 200)
	ErrDatabaseMigration = newErrcode("database migration error", 201)
	ErrDropProduction    = newErrcode("production database cannot be dropped", 202)

	// controllers errors (300-399)
	ErrInvalidParameters   = newErrcode("invalid parameter", 300)
	ErrNotFound            = newErrcode("not found", 301)
	ErrUnknown             = newErrcode("unknown", 302)
	ErrConfigurationFailed = newErrcode("configuration failed", 303)

	// auth errors (400-499)
	ErrUnauthorized = newErrcode("unauthorized", 400)
	ErrForbidden    = newErrcode("forbidden", 401)

	// business logic errors (500-599)
	ErrExternalLib        = newErrcode("external libraries error", 500)
	ErrSendingEmail       = newErrcode("email error", 501)
	ErrEmail              = newErrcode("error email", 501)
	ErrUserAlreadyExist   = newErrcode("user already exist", 502)
	ErrGeneratingToken    = newErrcode("error generation token", 503)
	ErrInvalidToken       = newErrcode("invalid token", 504)
	ErrTokenExpired       = newErrcode("invalid token", 505)
	ErrInvalidCredentials = newErrcode("invalid credentials", 506)
)

func newErrcode(message string, code int) GoCleanError {
	return GoCleanError{errors.New(message), code}
}

func Wrap(err *error, format string, args ...any) {
	if *err != nil {
		*err = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *err)
	}
}
