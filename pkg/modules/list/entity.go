package list

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
