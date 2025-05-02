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

	"github.com/QUDUSKUNLE/microservices/shared/logger"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/auth"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/organization"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/record"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"github.com/QUDUSKUNLE/microservices/shared/utils"
	"go.uber.org/zap"
)

func main() {
	if err := utils.LoadEnvironmentVariable(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize the logger
	logger.InitLogger()
	defer logger.Sync()

	// Initialize runtime server
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs("authorization", req.Header.Get("Authorization"))
		}),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register AuthServiceHandler
	if err := auth.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("AUTH"), opts); err != nil {
		log.Fatalf("Failed to register the auth service handler: %v", err)
	}

	// Register UserServiceHandler
	if err := user.RegisterUserServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("USER_SERVICE"), opts); err != nil {
		log.Fatalf("Failed to register the user service handler: %v", err)
	}

	// Register OrganizationServiceHandler
	if err := organization.RegisterOrganizationServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("ORGANIZATION"), opts); err != nil {
		log.Fatalf("Failed to register the organization service handler: %v", err)
	}

	// Register RecordServiceHandler
	if err := record.RegisterRecordServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("RECORD"), opts); err != nil {
		logger.GetLogger().Fatal("Failed to register the record service handler", zap.Error(err))
	}

	// Register DiagnosticServiceHandler
	if err := diagnostic.RegisterDiagnosticServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("DIAGNOSTIC"), opts); err != nil {
		logger.GetLogger().Fatal("Failed to register the diagnostic service handler", zap.Error(err))
	}

	// Register ScheduleServiceHandler
	if err := schedule.RegisterScheduleServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("SCHEDULE_SERVICE"), opts); err != nil {
		logger.GetLogger().Fatal("Failed to register the schedule service handler", zap.Error(err))
	}

	addr := fmt.Sprintf("%v:%v", os.Getenv("GATEWAY"), os.Getenv("GATEWAY_PORT"))
	go func() {
		logger.GetLogger().Info("Gateway server listening on port", zap.String("address", addr))
		if err := http.ListenAndServe(addr, mux); err != nil {
			logger.GetLogger().Fatal("Gateway server closed abruptly", zap.Error(err))
		}
	}()
	<-ctx.Done()

	logger.GetLogger().Info("Shutting down gateway server gracefully...")
	logger.GetLogger().Info("Gateway server stopped gracefully")
}
