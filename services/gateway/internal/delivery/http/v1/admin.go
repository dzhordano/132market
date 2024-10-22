package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initAdminRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// TODO Init all routes for manipulation with other services (orders, products, users, etc.)

	return r
}
