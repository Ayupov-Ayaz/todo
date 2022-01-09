package list

import (
	_errors "github.com/ayupov-ayaz/todo/errors"
	"github.com/ayupov-ayaz/todo/internal/models"
)

var ErrUpdateTodoListInvalid = _errors.BadRequest("update structure has not values")

type UpdateTodoList struct {
	Title       *string
	Description *string
}

func (u UpdateTodoList) Validate() error {
	if u.Title == nil && u.Description == nil {
		return ErrUpdateTodoListInvalid
	}

	return nil
}

type getAllListResponse struct {
	List []models.TodoList `json:"list"`
}
