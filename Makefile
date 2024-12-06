include local.env

LOCAL_BIN := $(CURDIR)/bin
LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN=${MIGRATION_DSN}

install golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... -- config .golangci.reference.yml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@v2.50.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.1.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
    GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generation-protoc:
	mkdir -p $(CURDIR)/pkg/chat_v1
	protoc --proto_path grpc/chat/v1 --proto_path vendor.protogen \
            --go_out=$(CURDIR)/pkg/chat_v1 --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=$(CURDIR)/bin/protoc-gen-go \
            --go-grpc_out=$(CURDIR)/pkg/chat_v1 --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=$(CURDIR)/bin/protoc-gen-go-grpc \
            --grpc-gateway_out=$(CURDIR)/pkg/chat_v1 --grpc-gateway_opt=paths=source_relative \
            --plugin=protoc-gen-grpc-gateway=$(CURDIR)/bin/protoc-gen-grpc-gateway \
            --validate_out=lang=go:$(CURDIR)/pkg/chat_v1 --validate_opt=paths=source_relative \
            --plugin=protoc-gen-validate=$(CURDIR)/bin/protoc-gen-validate \
            grpc/chat/v1/chat.proto

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi

migrate-new:
	mkdir -p migrations
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) create ${NAME} sql

migrate-up:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) up -v

migrate-down:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) down -v

migrate-reset:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) reset -v