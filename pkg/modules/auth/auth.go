package auth

import (
	"time"

	"github.com/ayupov-ayaz/todo/internal/models"
)

type UseCse interface {
	Create(user models.User) (int, error)
	SignIn(username, password string) (string, error)
}

type Repository interface {
	Create(user models.User) (int, error)
	Exist(username string) (bool, error)
	Get(username, password string) (models.User, error)
}

type CreateToken interface {
	CreateToken(userID int, lifeTime time.Duration) (string, error)
}
