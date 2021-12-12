package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultPort           = 8080
	defaultMaxHeaderBytes = 1 << 20 // 1MB
	defaultReadTimeout    = 10 * time.Second
	defaultWriteTimeout   = 10 * time.Second
)

type Cfg struct {
	Port           int
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func DefaultCfg() Cfg {
	return Cfg{
		Port:           defaultPort,
		MaxHeaderBytes: defaultMaxHeaderBytes,
		ReadTimeout:    defaultReadTimeout,
		WriteTimeout:   defaultWriteTimeout,
	}
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg Cfg) error {
	addr := ":" + strconv.Itoa(cfg.Port)

	s.httpServer = &http.Server{
		Addr:           addr,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}

	fmt.Println("listen and serve " + addr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
