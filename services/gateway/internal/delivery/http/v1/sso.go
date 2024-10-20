package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initSSORoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Post("/login", notImplemented)

	r.Post("/logout", notImplemented)

	r.Post("/register", notImplemented)

	// Обновление пары Access-Refresh токенов.
	r.Post("/refresh", notImplemented)

	// Запрос на смену пароля. Генерируется otp и отправляется на почту.
	r.Post("/reset-password", notImplemented)

	// Подтверждает смену пароля. После получения кода на почту заполняется тело типа:
	// reset_code: <code from email>
	// new_password: <new password>
	r.Post("/reset-password/confirm", notImplemented)

	// Проверяет токен на валидность. Если валиден - возвращает данные пользователя из токена. (Кста подумать, мб просто 200 кидать)
	r.Get("/verify-token", notImplemented)

	return r
}
