package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initUsersRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Post("/create", notImplemented)

	r.Put("/update/{userId}", notImplemented)

	r.Delete("/delete/{userId}", notImplemented)

	r.Get("/get/{userId}", notImplemented)

	// Получения информации о текущем пользователе (себе).
	r.Get("/me", notImplemented)

	// Только администраторам. Вроде...
	// Список всех пользователей.
	// Сделать пагинацию и т.д.
	// r.Get("/get", notImplemented)

	// Доступен только администраторам.
	// Изменение роли пользователя. Запрос:
	// role: <new role>
	r.Patch("/{userId}/role", notImplemented)

	// Доступен администраторам и модераторам.
	// Пользователь банится на определенный срок и прикрепляется причина блокировки. Запрос:
	// reason: <reason for block>
	// banned_at: <date>
	// banned_till: <date>
	// banned_by: <username>
	r.Patch("/{userId}/block", notImplemented)

	return r
}
