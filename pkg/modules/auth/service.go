package auth

import (
	"crypto/sha1"
	"fmt"

	_errors "github.com/ayupov-ayaz/todo/errors"

	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var (
	ErrUsernameIsBusy = _errors.BadRequest("username is busy")
)

type AuthorizationRepository interface {
	Create(user models.User) (int, error)
	Exist(username string) (bool, error)
}

type Service struct {
	repo     AuthorizationRepository
	validate *validator.Validate
	logger   *zap.Logger
	salt     []byte
}

func NewService(repo AuthorizationRepository, salt string) *Service {
	return &Service{
		repo:     repo,
		validate: validator.New(),
		logger:   zap.L().Named("auth_srv"),
		salt:     []byte(salt),
	}
}

func generatePasswordHash(pass string, salt []byte) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum(salt))
}

func (s *Service) Create(user models.User) (int, error) {
	if err := s.validate.Struct(user); err != nil {
		s.logger.Warn("validation user struct failed",
			zap.Error(err))

		return 0, err
	}

	exist, err := s.repo.Exist(user.Username)
	if err != nil {
		s.logger.Error("repo check user failed",
			zap.String("username", user.Username),
			zap.Error(err))

		return 0, err
	}

	if exist {
		s.logger.Warn("username is busy", zap.String("username", user.Username))
		return 0, ErrUsernameIsBusy
	}

	user.Password = generatePasswordHash(user.Password, s.salt)

	id, err := s.repo.Create(user)
	if err != nil {
		s.logger.Error("create user failed", zap.Error(err))
		return 0, err
	}

	return id, nil
}
