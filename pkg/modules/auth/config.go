package auth

import (
	"time"
)

type Config struct {
	Salt       []byte        `validate:"required"`
	SigningKey []byte        `validate:"required"`
	LifeTime   time.Duration `validate:"required,min=24h"`
}
