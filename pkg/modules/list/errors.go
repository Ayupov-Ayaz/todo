package list

import _errors "github.com/ayupov-ayaz/todo/errors"

var (
	ErrInvalidRequest        = _errors.BadRequest("invalid request")
	ErrUpdateTodoListInvalid = _errors.BadRequest("update structure has not values")
	ErrListNotFound          = _errors.NotFound("list not found or list does not belong to you")
)
