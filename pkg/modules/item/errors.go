package item

import _errors "github.com/ayupov-ayaz/todo/internal/errors"

var (
	ErrItemNotFound          = _errors.NotFound("item not found")
	ErrListDoesntBelongsUser = _errors.Forbidden("list doesn't belong user")
	ErrItemDoesntBelongUser  = _errors.Forbidden("item doesn't belong user")
)
