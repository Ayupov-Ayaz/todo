package auth

import (
	"errors"
	"time"
)

type Config struct {
	Salt       []byte
	SigningKey []byte
	LifeTime   time.Duration
}

func (c Config) Validate() error {
	if len(c.Salt) == 0 {
		return errors.New("salt is required")
	}

	if len(c.SigningKey) == 0 {
		return errors.New("signing key is required")
	}

	if c.LifeTime == 0 {
		return errors.New("lifetime is required")
	}

	return nil
}
