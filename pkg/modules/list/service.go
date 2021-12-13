package list

type TodoListRepository interface {
}

type Service struct {
	repo TodoListRepository
}

func NewService(repo TodoListRepository) *Service {
	return &Service{repo: repo}
}
