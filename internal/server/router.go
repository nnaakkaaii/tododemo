package server

import (
	"context"
	"fmt"
	"github.com/nnaakkaaii/tododemo/internal/db"
	"github.com/nnaakkaaii/tododemo/internal/handler"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(port int, d db.DB) *Server {
	mux := http.NewServeMux()
	mux.Handle("/todos", handler.NewTODOsHandler(d))
	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to server: %v", err)
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown: %v", err)
	}
	return nil
}
