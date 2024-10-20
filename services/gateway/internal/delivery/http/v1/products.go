package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initProductsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Создание продукта на маркетплейсе.
	// Такой продукт следует модерировать перед публикацией.
	r.Post("/create", notImplemented)

	// Получение всех продуктов.
	// Пагинация, категория, теги и т.д. должны быть в параметрах запроса.
	r.Get("/search", notImplemented)

	// Продукты, которые принадлежат пользователю.
	// Также возвращается статус продукта, просмотры и т.п.
	r.Get("/my", notImplemented)

	r.Get("/{productId}", notImplemented)

	// Думаю также после обновления есть "проблема" в модерации.
	r.Put("/{productId}", notImplemented)

	r.Delete("/{productId}", notImplemented)

	return r
}
