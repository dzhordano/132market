package http

import (
	v1 "dzhordano/132market/services/gateway/internal/delivery/http/v1"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h *handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	handlerV1 := v1.New()

	r.Mount("/api", handlerV1.InitRoutes())

	// TODO Init all other routes

	return r
}
