include .env
MIGRATIONS_DSN="host=${POSTGRES_HOST} port=${POSTGRES_PORT} dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=${POSTGRES_SSLMODE}"
LOCAL_BIN=$(CURDIR)/.bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-users-api

# Consider changing --plugin flags suffixes if not on windows; dunno if it'll work with .exe if not using on windows :P.
generate-users-api:
	mkdir -p pkg/users_v1
	protoc --proto_path internal/proto/users_v1 \
	--go_out=pkg/users_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=.bin/protoc-gen-go.exe \
	--go-grpc_out=pkg/users_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=.bin/protoc-gen-go-grpc.exe \
	internal/proto/users_v1/users.proto

init-db:
	docker run --name=132mp-users-db -e POSTGRES_PASSWORD=postgres -p 5433:5432 -d postgres
	$(MAKE) migrate

migrate:
	goose -dir ./migrations postgres $(MIGRATIONS_DSN) up