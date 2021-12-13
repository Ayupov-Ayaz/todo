package auth

type AuthorizationRepository interface {
}

type Service struct {
	repo AuthorizationRepository
}

func NewService(repo AuthorizationRepository) *Service {
	return &Service{repo: repo}
}
