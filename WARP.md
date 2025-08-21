# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Common Development Commands

### Building and Running
```bash
# Build all packages in the repository
go build -v ./...

# Run a specific package (e.g., the load balancer)
go run loadbalancer/main.go

# Build with output binary
go build -o bin/loadbalancer loadbalancer/main.go
```

### Testing
```bash
# Run all tests
go test -v ./...

# Run tests for a specific package
go test -v ./arrays

# Run a single test function
go test -v -run TestSum ./arrays

# Run tests with coverage
go test -cover ./...

# Run tests with coverage profile
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Module Management
```bash
# Download dependencies
go mod download

# Tidy up go.mod and go.sum
go mod tidy

# Verify dependencies
go mod verify

# View dependency graph
go mod graph
```

### Code Quality
```bash
# Format all Go files
go fmt ./...

# Run static analysis
go vet ./...

# Check for potential issues
go test -race ./...
```

## Project Architecture

This repository is a **Go learning and tutorial collection** organized into themed packages. Each directory represents a specific Go concept or implementation pattern.

### Core Structure

- **Educational Modules**: Each top-level directory focuses on a specific Go concept:
  - `arrays/` - Array and slice operations with comprehensive tests
  - `channels/` - Goroutine communication patterns, worker pools, and channel management
  - `Algorithm/` - Common algorithms (binary search, quick sort, linear search)
  - `DesignPatterns/` - GoF design patterns (Facade, Adapter, Singleton)
  - `loadbalancer/` - Production-quality HTTP load balancer implementation

- **Advanced Topics**: Complex real-world implementations:
  - `Artificial/` - Separate module with its own go.mod for AI/ML experiments
  - `kafka/` - Kafka integration patterns (separate module)
  - `DesignPatterns/Facade/Shipping/` - Hexagonal architecture implementation

### Key Dependencies

- **Core**: Uses Go 1.23.0+ with toolchain 1.24.3
- **Protocol Buffers**: `google.golang.org/protobuf` for serialization
- **gRPC**: `google.golang.org/grpc` for RPC communication  
- **Telegram Bot API**: `github.com/go-telegram-bot-api/telegram-bot-api/v5`
- **Text Processing**: `golang.org/x/text` for internationalization

### Testing Philosophy

The codebase follows test-driven development with comprehensive test coverage:
- Each package includes `*_test.go` files
- Tests use standard Go testing patterns with table-driven tests
- Example: `arrays/arrays_test.go` tests all array manipulation functions

### Load Balancer Architecture

The `loadbalancer/main.go` is a production-ready HTTP load balancer featuring:
- **Balancing Policies**: Round-robin, least connections, random
- **Health Checking**: Periodic backend health monitoring with configurable intervals
- **Sticky Sessions**: Cookie-based client-to-backend affinity
- **Retry Logic**: Automatic failover with configurable retry attempts
- **Graceful Shutdown**: Signal handling with context cancellation

Key components:
- `Backend` struct: Manages individual upstream servers with atomic counters
- `LoadBalancer` struct: Orchestrates traffic distribution and health checks
- Supports weighted backends via CLI arguments (e.g., `http://server:port|weight`)

### Design Patterns Implementation

The `DesignPatterns/` directory demonstrates:
- **Facade Pattern**: Simplifies complex subsystem interactions (Order processing)
- **Hexagonal Architecture**: Clean separation in the Shipping service with:
  - `core/domain/` - Business logic and entities
  - `adapters/` - External system integrations (external_adapter, internal_adapter)
  - `ports/` - Interface definitions

### Module Organization

The repository uses Go modules effectively:
- Root module: `github.com/QUDUSKUNLE/Golang/tutorial`
- Independent modules: `Artificial/` and `kafka/` have separate go.mod files
- This allows isolated dependency management for different concerns

### Development Workflow

1. **Adding New Concepts**: Create a new directory with focused Go files
2. **Testing**: Always include comprehensive tests following existing patterns
3. **Documentation**: Use detailed comments explaining concepts (see `loadbalancer/main.go`)
4. **Dependencies**: Run `go mod tidy` after adding new imports
