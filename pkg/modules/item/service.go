package item

type TodoItemRepository interface {
}

type Service struct {
	repo TodoItemRepository
}

func NewService(repo TodoItemRepository) *Service {
	return &Service{repo: repo}
}
