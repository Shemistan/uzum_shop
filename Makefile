test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

lint:
	golangci-lint run

BIN:=$(CURDIR)/bin
install-deps:
	GOBIN=$(BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2

gen: ## Генерация proto-файлов
		mkdir -p pkg/shop_v1
		protoc --proto_path=api/shop_v1 \
        				--proto_path=proto \
        				--go_out=pkg/shop_v1 --go_opt=paths=source_relative \
        				--plugin=protoc-gen-go=bin/protoc-gen-go \
        				--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
        				--go-grpc_out=pkg/shop_v1 --go-grpc_opt=paths=source_relative \
        				--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
        				--grpc-gateway_out=pkg/shop_v1 --grpc-gateway_opt=paths=source_relative \
        				api/shop_v1/shop.proto