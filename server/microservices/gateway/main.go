package main

// Package main serves as the entry point for the gRPC Gateway server.
// It sets up HTTP handlers to proxy requests to gRPC services and provides
// RESTful endpoints for the microservices in the application.
//
// This file imports the necessary packages for setting up the gRPC Gateway,
// including the gRPC runtime, credentials, and metadata handling. It also
// imports the generated protocol buffer code for the various microservices
// (organization, record, user, and auth) and utility functions.
//
// The gRPC Gateway enables seamless communication between RESTful clients
// and gRPC services by translating HTTP/JSON requests into gRPC requests
// and vice versa.
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/QUDUSKUNLE/microservices/gateway/constants"
	"github.com/QUDUSKUNLE/microservices/gateway/middleware"
	"github.com/QUDUSKUNLE/microservices/notification-service/subscribe"
	"github.com/QUDUSKUNLE/microservices/shared/events"
	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/auth"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/record"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"go.uber.org/zap"
)

const (
	rateLimitRequestsPerSecond = 10
	rateLimitBurst             = 5
	maxRequestBodySize         = 10 * 1024 * 1024 // 10 MB
)

func main() {
	cfg, err := utils.LoadConfig("GATEWAY")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize runtime server
	mux := setupRuntimeServer()

	// Register services
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	registerServices(ctx, mux, opts, cfg)

	// Create a  rate limiter middleware
	handler := setupMiddleware(mux)

	addr := fmt.Sprintf("%v:%v", cfg.Gateway, cfg.Port)

	go events.EventsSubscriber(
		os.Getenv("KAFKA_BROKER"),
		os.Getenv("KAFKA_TOPIC"),
		os.Getenv("KAFKA_GROUP_ID"),
		subscribe.SubsribeNotification,
	)

	go startHTTPServer(addr, handler)
	<-ctx.Done()
	logger.GetLogger().Info("Shutting down gateway server gracefully...")
}

func registerServices(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption, cfg *utils.Config) {
	services := map[string]func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		constants.AuthService:       auth.RegisterAuthServiceHandlerFromEndpoint,
		constants.UserService:       user.RegisterUserServiceHandlerFromEndpoint,
		constants.RecordService:     record.RegisterRecordServiceHandlerFromEndpoint,
		constants.DiagnosticService: diagnostic.RegisterDiagnosticServiceHandlerFromEndpoint,
		constants.ScheduleService:   schedule.RegisterScheduleServiceHandlerFromEndpoint,
	}
	for serviceName, registerFunc := range services {
		var endpoint string
		switch serviceName {
		case constants.AuthService:
			endpoint = cfg.AuthService
		case constants.UserService:
			endpoint = cfg.UserService
		case constants.RecordService:
			endpoint = cfg.RecordService
		case constants.DiagnosticService:
			endpoint = cfg.DiagnosticService
		case constants.ScheduleService:
			endpoint = cfg.ScheduleService
		default:
			logger.GetLogger().Fatal(fmt.Sprintf("Unknown service: %s", serviceName))
		}
		if err := registerFunc(ctx, mux, endpoint, opts); err != nil {
			logger.GetLogger().Fatal(fmt.Sprintf("Failed to register %s handler", serviceName), zap.Error(err))
		}
		logger.GetLogger().Info(fmt.Sprintf("%s registered successfully", serviceName))
	}
}

func startHTTPServer(addr string, handler http.Handler) {
	logger.GetLogger().Info("Gateway server listening", zap.String("address", addr))
	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.GetLogger().Fatal("Gateway server closed abruptly", zap.Error(err))
	}
}

// setupRuntimeServer initializes the gRPC Gateway runtime server.
func setupRuntimeServer() *runtime.ServeMux {
	return runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs("authorization", req.Header.Get("Authorization"))
		}),
	)
}

// setupMiddleware sets up the middleware pipeline.
func setupMiddleware(mux *runtime.ServeMux) http.Handler {
	rateLimiter := middleware.NewRateLimiter(rateLimitRequestsPerSecond, rateLimitBurst)
	rateLimitedMux := middleware.RateLimitMiddleware(rateLimiter, mux)
	corsMux := middleware.CORSMiddleware(rateLimitedMux)
	return middleware.LimitRequestBodyMiddleware(maxRequestBodySize, corsMux)
}
