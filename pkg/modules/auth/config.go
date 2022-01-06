package auth

import (
	"errors"
	"time"
)

type Config struct {
	Salt       []byte
	SigningKey []byte
	Expired    time.Duration
}

func (c Config) Validate() error {
	if len(c.Salt) == 0 {
		return errors.New("salt is required")
	}

	if len(c.SigningKey) == 0 {
		return errors.New("signing key is required")
	}

	if c.Expired == 0 {
		return errors.New("expired is required")
	}

	return nil
}
