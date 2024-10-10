package main

import (
	"context"
	"dzhordano/132market/services/users/internal/app"
)

// TODO Шаги к доработке:
// Доделать круд лол
// Обертки для ошибок + обертка БД
// Доступ сюда получают только админы ЛИБО сюда приходит SSO-сервис с клиента при регистрации/логине
// Метрики, алерты, паттерны безопасности, логирование и тд.
// Написать тесты
// Написать доку

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		panic(err)
	}

	if err = a.Run(ctx); err != nil {
		panic(err)
	}
}
