package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

func NewServer(port string, r *gin.Engine) *Server {
	return &Server{
		server: &http.Server{
			Addr:    ":"+port,
			Handler: r,
		},
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown: %w", err)
	}

	return nil
}
