package auth

import _errors "github.com/ayupov-ayaz/todo/internal/errors"

var (
	ErrInvalidRequest      = _errors.BadRequest("invalid request")
	ErrAuthorizationFailed = _errors.Forbidden("authorization failed")
	ErrUsernameIsBusy      = _errors.BadRequest("username is busy")
)
