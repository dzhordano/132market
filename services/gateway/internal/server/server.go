package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	httpServer *http.Server
}

func New(address string, h http.Handler) *HttpServer {
	return &HttpServer{
		httpServer: &http.Server{
			Addr:           address,
			Handler:        h,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   5 * time.Second,
			IdleTimeout:    120 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *HttpServer) Run() error {
	log.Printf("Starting http server on address: %s", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *HttpServer) Shutdown() error {
	log.Printf("Shutting down http server")

	return s.httpServer.Shutdown(context.TODO())
}
