GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

.PHONY: setup generate migrate-up migrate-down protoc run_local_server generate-tls-cert run-tls-server test-tls build-gateway run-gateway test-full-stack

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
	--go_out=../protogen \
 	--go-grpc_opt=paths=source_relative \
	--go-grpc_out=../protogen \
	--grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_out=../protogen \
	--grpc-gateway_opt generate_unbound_methods=true \
	./**/*.proto

run_local_server:
	air

# TLS certificate generation for development
# This creates a self-signed certificate valid for 365 days
# Note: For production, use properly signed certificates from a trusted CA
generate-tls-cert:
	@echo "Generating self-signed TLS certificates for development..."
	@mkdir -p certs
	@openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
		-keyout certs/server.key -out certs/server.crt \
		-subj "/CN=localhost" \
		-addext "subjectAltName = DNS:localhost,IP:127.0.0.1"
	@echo "Certificates generated in the certs directory"
	@echo "  - Private key: certs/server.key"
	@echo "  - Certificate: certs/server.crt"

# Run the server with TLS enabled
# This builds and runs the server with the TLS configuration
run-tls-server: generate-tls-cert
	@echo "Building and running server with TLS enabled..."
	@go build -o bin/record-service main.go
	@./bin/record-service

# Test the TLS configuration
# This ensures certificates exist and runs the TLS tests
test-tls:
	@echo "Testing TLS configuration..."
	@if [ ! -f certs/server.crt ] || [ ! -f certs/server.key ]; then \
		echo "TLS certificates not found. Generating..."; \
		$(MAKE) generate-tls-cert; \
	fi
	@echo "Starting TLS server in background for testing..."
	@go build -o bin/record-service main.go
	@./bin/record-service > /dev/null 2>&1 & echo $$! > .server.pid
	@echo "Waiting for server to start..."
	@sleep 2
	@echo "Running TLS tests..."
	@go test -v ./test/tls_test.go || { echo "Tests failed!"; kill `cat .server.pid`; rm .server.pid; exit 1; }
	@echo "✅ TLS tests passed!"
	@kill `cat .server.pid` && rm .server.pid
	@echo "Test server stopped."

# Build the API gateway
# This compiles the gateway binary
build-gateway: generate-tls-cert
	@echo "Building Gateway server..."
	@mkdir -p bin
	@go build -o bin/gateway cmd/gateway/main.go
	@echo "Gateway binary built at bin/gateway"

# Run the gateway with TLS
# This starts the gateway server with TLS enabled
run-gateway: build-gateway
	@echo "Starting Gateway server with TLS enabled..."
	@./bin/gateway --grpc-server-endpoint=localhost:3012 --gateway-port=7556

# Test the full stack (gRPC + Gateway)
# This runs both the gRPC server and gateway together
test-full-stack: generate-tls-cert
	@echo "Testing full stack (gRPC + Gateway)..."
	@echo "Building servers..."
	@go build -o bin/record-service main.go
	@go build -o bin/gateway cmd/gateway/main.go
	
	@echo "Starting gRPC server in background..."
	@./bin/record-service > /dev/null 2>&1 & echo $$! > .grpc.pid
	@echo "Waiting for gRPC server to start..."
	@sleep 2
	
	@echo "Starting Gateway server in background..."
	@./bin/gateway > /dev/null 2>&1 & echo $$! > .gateway.pid
	@echo "Waiting for Gateway server to start..."
	@sleep 2
	
	@echo "Running API test..."
	@curl -k -s -o /dev/null -w "HTTP Status Code: %{http_code}\n" \
		-X GET https://localhost:7556/v1/records || \
		{ echo "Failed to connect to Gateway"; \
		kill `cat .grpc.pid` 2>/dev/null; rm .grpc.pid 2>/dev/null; \
		kill `cat .gateway.pid` 2>/dev/null; rm .gateway.pid 2>/dev/null; exit 1; }
	
	@echo "Stopping servers..."
	@kill `cat .grpc.pid` 2>/dev/null; rm .grpc.pid 2>/dev/null
	@kill `cat .gateway.pid` 2>/dev/null; rm .gateway.pid 2>/dev/null
	@echo "✅ Full stack test completed successfully!"
