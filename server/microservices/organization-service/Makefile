GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

.PHONY: setup generate migrate-up migrate-down protoc run_local_server

SQLC_VERSION=v1.27.0
MIGRATE_VERSION=v4.17.0

setup:
	go mod download
	GOBIN=$(GOBIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@$(SQLC_VERSION)
	GOBIN=$(GOBIN) go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)

generate:
	$(GOBIN)/sqlc generate -f db/sqlc.json

migrate-up:
	$(GOBIN)/migrate -path="db/migrations" -database "postgres://abumuhsinah:abumuhsinah@localhost:5432/golang?sslmode=disable" up

migrate-down:
	$(GOBIN)/migrate -path="db/migrations" -database "postgres://abumuhsinah:abumuhsinah@localhost:5432/golang?sslmode=disable" down

protoc:
	cd proto && protoc --go_opt=paths=source_relative \
	--go_out=../protogen/golang \
 	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang \
	--grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/golang \
	--grpc-gateway_opt generate_unbound_methods=true \
	./**/*.proto

run_local_server:
	air
