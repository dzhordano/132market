package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initChatRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Создание сессии между пользователями после.
	// Свойства чата:
	// Соединение будет поддерживаться 72 часа после последнего сообщения, после чего сессия будет удалена.
	r.Post("/create", notImplemented)

	// Подключение к чату. Это уже получается WebSocket соединение (а в gRPC stream, полагаю).
	// с ID идея такая: user1_uuid + user2_uuid = chat_uuid. TODO надо вариант получше.
	r.Post("/connect/{chatId}", notImplemented)

	r.Post("/send", notImplemented)

	// Получении истории чата.
	r.Get("/{chatId}", notImplemented)

	return r
}
