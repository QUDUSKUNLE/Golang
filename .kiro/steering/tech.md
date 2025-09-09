# Technology Stack

## Core Technologies
- **Go**: Version 1.23.0+ (with toolchain 1.24.3)
- **gRPC**: For microservice communication
- **Protocol Buffers**: Service interface definitions
- **Docker & Docker Compose**: Containerization and orchestration

## Web Frameworks
- **Fiber**: High-performance web framework (used in goFiber server)
- **Echo**: Web framework (used in goEcho server)

## Key Dependencies
- `google.golang.org/protobuf`: Protocol buffer support
- `google.golang.org/grpc`: gRPC framework
- `github.com/go-telegram-bot-api/telegram-bot-api/v5`: Telegram bot integration
- `golang.org/x/text`: Text processing utilities

## Development Tools
- **Air**: Live reload for Go applications (`.air.toml` configs present)
- **SQLC**: SQL code generation
- **Swagger**: API documentation
- **Trunk**: Code quality and linting

## Infrastructure
- **Kafka**: Event streaming and messaging
- **PostgreSQL**: Database (inferred from migration files)
- **Redis**: Caching (inferred from cache scripts)

## Common Commands

### Building and Running
```bash
# Build main application
go build -o main main.go

# Run with arguments (for video processing)
./main <YouTube-URL>

# Run tests
go test ./...

# Run specific package tests
go test ./Hello

# Run with coverage
go test -cover ./...
```

### Development
```bash
# Live reload (in directories with .air.toml)
air

# Generate protobuf code
./protocomp.sh

# Generate Swagger docs
./swag.sh

# Database migrations (in microservices)
make migrate-up
make migrate-down
```

### Docker Operations
```bash
# Build and run all services
docker-compose up --build

# Run in development mode
docker-compose -f docker-compose.dev.yml up

# Stop services
docker-compose down

# View logs
docker logs <container-name>
```

### Testing
```bash
# Run test scripts
./test.sh

# Run specific service tests
./Server/goFiber/test.sh
./Server/goEcho/test.sh
```

## Project Module
- Module path: `github.com/QUDUSKUNLE/Golang/tutorial`
- Import pattern: Use relative imports within the module
