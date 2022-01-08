package auth

import (
	"crypto/sha1"
	"fmt"
	"time"

	_errors "github.com/ayupov-ayaz/todo/errors"

	"github.com/ayupov-ayaz/todo/internal/models"
	"github.com/ayupov-ayaz/todo/pkg/services/validator"
	"go.uber.org/zap"
)

var (
	ErrUsernameIsBusy = _errors.BadRequest("username is busy")
)

type AuthorizationRepository interface {
	Create(user models.User) (int, error)
	Exist(username string) (bool, error)
	Get(username, password string) (models.User, error)
}

type CreateToken interface {
	CreateToken(userID int, lifeTime time.Duration) (string, error)
}

type Service struct {
	repo     AuthorizationRepository
	validate validator.Validator
	token    CreateToken
	salt     []byte
	lifeTime time.Duration
	logger   *zap.Logger
}

func NewService(repo AuthorizationRepository, val validator.Validator, token CreateToken, salt []byte, lifeTime time.Duration) *Service {
	return &Service{
		repo:     repo,
		validate: val,
		token:    token,
		salt:     salt,
		lifeTime: lifeTime,
		logger:   zap.L().Named("auth_srv"),
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

func (s *Service) SignIn(username, password string) (string, error) {
	user, err := s.repo.Get(username, generatePasswordHash(password, s.salt))
	if err != nil {
		s.logger.Error("get user failed",
			zap.String("username", username),
			zap.Error(err))

		return "", err
	}

	token, err := s.token.CreateToken(user.ID, s.lifeTime)
	if err != nil {
		s.logger.Error("create jwt token failed", zap.Error(err))

		return "", err
	}

	return token, nil
}
