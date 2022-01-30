package jwt

import (
	"time"

	_errors "github.com/ayupov-ayaz/todo/internal/errors"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type Service struct {
	signingKey []byte
}

func NewUseCase(signingKey []byte) Service {
	return Service{signingKey: signingKey}
}

func (s Service) CreateToken(userID int, lifeTime time.Duration) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(lifeTime).Unix(),
		},
		UserID: userID,
	})

	jwtToken, err := token.SignedString(s.signingKey)
	if err != nil {

		return "", err
	}

	return jwtToken, nil
}

func (s Service) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, _errors.Forbidden("invalid signing method")
		}

		return s.signingKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, _errors.Forbidden("invalid claims type")
	}

	return claims.UserID, nil
}
