package item

type UpdateItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (u UpdateItem) Validate() error {
	if u.Title == nil && u.Description == nil && u.Done == nil {
		return ErrUpdateItemInvalid
	}

	return nil
}
