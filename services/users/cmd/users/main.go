package main

import (
	"dzhordano/132market/services/users/config"
	"dzhordano/132market/services/users/internal/application/services"
	"dzhordano/132market/services/users/internal/infrastructure/db/postgres"
	"dzhordano/132market/services/users/internal/infrastructure/grpc"
	"dzhordano/132market/services/users/pkg/logger"
	"log/slog"
	"os"
)

// TODO
// - (-)Добавить фильтр поиска (тип только активные, заблокированные и т.д.)
// - (+)Конфигурация + (-)graceful shutdown
// - (+)Обернуть ошибки + намутить хороший вывод для gRPC. (-)Не забыть про возврат множественных ошибок (например при валидации: возврат перечисления конкретных полей)
// - Разрулить то, как я буду эти же ошибки потом выплевывать на Gateway
// - (-)Обертка над БД
// - (-)Логирование ELK
// - (-)DI-контейнер
// - (+)Unit + (-)Integration + (-)Нагрузочные тесты
// - (-)Собрать как докер контейнер
// - (-)Собрать пайплайн
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
