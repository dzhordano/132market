package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initOrdersRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Post("/create", notImplemented)

	r.Get("/{orderId}", notImplemented)

	// Получения заказов пользователя.
	r.Get("/my", notImplemented)

	r.Delete("/{orderId}", notImplemented)

	r.Patch("/{orderId}/status", notImplemented)

	return r
}
