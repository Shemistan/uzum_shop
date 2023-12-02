ifneq (,$(wildcard ./build/.env))
    include ./build/.env
    export $(shell sed 's/=.*//' ./build/.env)
endif

gen_proto:gen_login gen_shop

gen_shop:
		@mkdir -p ./docs
		@mkdir -p pkg/shopV1
		@protoc	--go_out=pkg/shopV1 --go_opt=paths=source_relative \
				--go-grpc_out=pkg/shopV1 --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=pkg/shopV1 --grpc-gateway_opt=paths=source_relative \
				--openapiv2_out=allow_merge=true,merge_file_name=api_shopV1:docs \
				internal/protos/shop.proto

gen_login:
		@mkdir -p pkg/loginV1
		@protoc 	--go_out=pkg/loginV1 --go_opt=paths=source_relative \
				--go-grpc_out=pkg/loginV1 --go-grpc_opt=paths=source_relative \
				internal/protos/login.proto

docker_create_network:
	@docker network create uzum-api || true

docker_infra: docker_create_network
	@docker-compose -f ./build/docker_compose_infra.yml -p uzum up -d

migration_up:
	@goose -dir ./migrations postgres "${DBCONNSTR}" up

migration_down:
	@goose -dir ./migrations postgres "${DBCONNSTR}" down

migration_create:
	@goose -dir ./migrations postgres "${DBCONNSTR}" create default sql

run:
	@go run cmd/main.go