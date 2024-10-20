package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initStatisticsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Получение всей статистики по продуктам.
	// Только админы.
	r.Get("/products", notImplemented)

	// Получение всей статистики по заказам пользователя.
	// Только админы.
	r.Get("/orders", notImplemented)

	// Получение всей статистики по пользователям.
	// Может получить также и пользователь.
	r.Get("/users/{userId}", notImplemented)

	return r
}
