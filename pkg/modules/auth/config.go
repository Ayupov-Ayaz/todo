package auth

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Salt       []byte        `validate:"required"`
	SigningKey []byte        `validate:"required"`
	LifeTime   time.Duration `validate:"required,min=24h"`
}

func InitConfig() Config {
	return Config{
		Salt:       []byte(os.Getenv("PASS_SALT")),
		SigningKey: []byte(os.Getenv("SIGNING_KEY")),
		LifeTime:   viper.GetDuration("auth.lifetime"),
	}
}
