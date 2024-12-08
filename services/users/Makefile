LOCAL_BIN:=$(CURDIR)/.bin


install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api-v1

generate-user-api-v1:
	mkdir -p pkg/pb
	mkdir -p pkg/pb/user_v1
	protoc --proto_path api/proto/user_v1 \
	--go_out=pkg/pb/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=.bin/protoc-gen-go.exe \
	--go-grpc_out=pkg/pb/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=.bin/protoc-gen-go-grpc.exe \
	api/proto/user_v1/user.proto

include .env

init-db:
	docker run --name=users-db -p ${POSTGRES_PORT}:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres

test.integration:
	make init-test-db
	go test -v ./tests/integration
	docker stop users-test-db
	docker rm users-test-db

init-test-db:
	docker run --name=users-test-db -p ${POSTGRES_TEST_PORT}:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres

goose-up:
	goose -dir migrations postgres ${POSTGRES_DSN} up

mockgen:
	mockgen -source=internal/application/interfaces/user_service.go -destination=internal/interfaces/grpc/mocks/mocks.go 

ghz.findbyid.success:
	ghz \
		--proto api/proto/user_v1/user.proto \
		--call api.UserServiceV1.FindUserById \
		--data '{"id": "52a82ad8-b265-11ef-874c-0250e2d265f1"}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:55001

ghz.findbycreds.success:
	ghz \
		--proto api/proto/user_v1/user.proto \
		--call api.UserServiceV1.FindUserByCredentials \
		--data '{"email": "admin@null.ru", "password": "null_password"}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:55001


ghz.findbyid.error:
	ghz \
		--proto api/proto/user_v1/user.proto \
		--call api.UserServiceV1.FindUserById \
		--data '{"id": "00000000-0000-0000-0000-000000000000"}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:55001