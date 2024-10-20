package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initMailingRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Отправка кода для подтверждения почты.
	r.Post("/send", notImplemented)

	// Подтверждение почты.
	r.Post("/verify/{code}", notImplemented)

	// TODO Подумать про уведомления...
	// Проблемы:
	// - Если отправлять уведомления, то кому (как хранить инфу о тех, кому отправлять уведомления?).
	// - Где хранить подписчиков? (думаю нужна nosql бд).

	return r
}
