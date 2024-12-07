package main

import (
	"log/slog"
	"os"

	"github.com/dzhordano/132market/services/users/config"
	"github.com/dzhordano/132market/services/users/internal/application/services"
	"github.com/dzhordano/132market/services/users/internal/infrastructure/db/postgres"
	"github.com/dzhordano/132market/services/users/internal/infrastructure/grpc"
	"github.com/dzhordano/132market/services/users/pkg/logger"
)

// TODO
// - (-)Добавить фильтр поиска (тип только активные, заблокированные и т.д.) [admin]
// - (-)graceful shutdown
// - (+)Обернуть ошибки + намутить хороший вывод для gRPC. (-)Не забыть про возврат множественных ошибок
// (например при валидации: возврат перечисления конкретных полей) {errors.Join(args ...error) error юзать}
// - Разрулить то, как я буду эти же ошибки потом выплевывать на Gateway {буду преобразовывать просто на самом gateway}
// - (-)Обертка над БД
// - (-)Логирование ELK
// - (-)DI-контейнер
// - (-)Нагрузочные тесты
// - (-)Рефакторинг кода
// - (-)Разобраться с CQRS либо убрать вообще разделение
// TODO После поднятия SSO сервиса
// - (-)Поднять клиент к SSO Token Validator API
// - (-)Использовать TLS

func main() {
	logger := logger.NewTintSlogLogger(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // TODO не над хардкодить левел...
	})

	pool := postgres.NewPool(config.MustNewPostgresConfig().DSN())

	repo := postgres.NewUserRepository(pool)

	svc := services.NewUserService(logger, repo)

	serv := grpc.NewServer(svc)

	logger.Info("Starting server...")

	if err := serv.Run(config.MustNewGrpcConfig().Address()); err != nil {
		slog.Error("Failed to run server:", slog.String("error", err.Error()))
	}

}
