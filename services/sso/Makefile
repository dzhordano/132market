LOCAL_BIN:=$(CURDIR)/.bin


install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-sso-api-v1

generate-sso-api-v1:
	mkdir -p pkg/pb
	mkdir -p pkg/pb/sso_v1
	protoc --proto_path api/proto/sso_v1 \
	--go_out=pkg/pb/sso_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=.bin/protoc-gen-go.exe \
	--go-grpc_out=pkg/pb/sso_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=.bin/protoc-gen-go-grpc.exe \
	api/proto/sso_v1/sso.proto

include .env

init-db:
	docker run --name=sso-db -p ${POSTGRES_PORT}:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres

test.integration:
	make init-test-db
	go test -v ./tests/integration
	docker stop sso-test-db
	docker rm sso-test-db

init-test-db:
	docker run --name=sso-test-db -p ${POSTGRES_TEST_PORT}:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres

goose-up:
	goose -dir migrations postgres ${POSTGRES_DSN} up

mockgen:
	mockgen -source=internal/application/interfaces/sso_service.go -destination=internal/interfaces/grpc/mocks/mocks.go 