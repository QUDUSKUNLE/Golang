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

	"github.com/QUDUSKUNLE/microservices/gateway/middleware"
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
	if err := utils.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize runtime server
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs("authorization", req.Header.Get("Authorization"))
		}),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	registerServices(ctx, mux, opts)

	// Create a  rate limiter middleware
	rateLimiter := middleware.NewRateLimiter(rateLimitRequestsPerSecond, rateLimitBurst)
	rateLimitedMux := middleware.RateLimitMiddleware(rateLimiter, mux)
	corsMux := middleware.CORSMiddleware(rateLimitedMux)
	limitRequestBodyMux := middleware.LimitRequestBodyMiddleware(maxRequestBodySize, corsMux)

	addr := fmt.Sprintf("%v:%v", os.Getenv("GATEWAY"), os.Getenv("GATEWAY_PORT"))
	go startHTTPServer(addr, limitRequestBodyMux)
	<-ctx.Done()
	logger.GetLogger().Info("Shutting down gateway server gracefully...")
}

func registerServices(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) {
	services := map[string]func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		"AuthService":       auth.RegisterAuthServiceHandlerFromEndpoint,
		"UserService":       user.RegisterUserServiceHandlerFromEndpoint,
		"RecordService":     record.RegisterRecordServiceHandlerFromEndpoint,
		"DiagnosticService": diagnostic.RegisterDiagnosticServiceHandlerFromEndpoint,
		"ScheduleService":   schedule.RegisterScheduleServiceHandlerFromEndpoint,
	}
	for serviceName, registerFunc := range services {
		endpoint := os.Getenv(serviceName)
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
