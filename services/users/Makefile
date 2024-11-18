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
	protoc --proto_path api/proto/user_v1 \
	--go_out=pkg/pb/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=.bin/protoc-gen-go.exe \
	--go-grpc_out=pkg/pb/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=.bin/protoc-gen-go-grpc.exe \
	api/proto/user_v1/user.proto

include .env
init-db:
	docker run --name=users-db -p ${POSTGRES_PORT}:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres

goose-up:
	goose -dir migrations postgres ${POSTGRES_DSN} up