package internal

import (
	"context"
	"net/http"
	

	config "github.com/kovalyov-valentin/orders-service/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.HTTPServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.Address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
		IdleTimeout:    cfg.IdleTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
