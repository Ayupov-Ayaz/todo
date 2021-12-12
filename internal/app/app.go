package app

import (
	"fmt"

	"github.com/ayupov-ayaz/todo/internal/delivery/http"
)

func Run() error {
	s := new(http.Server)
	if err := s.Run(http.DefaultCfg()); err != nil {
		return fmt.Errorf("occured while running http server: %w", err)
	}

	return nil
}
