package server

import (
	"github.com/QUDUSKUNLE/microservices/schedule-service/adapters/handler"
	"github.com/QUDUSKUNLE/microservices/schedule-service/adapters/repository"
	"github.com/QUDUSKUNLE/microservices/schedule-service/core/services"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/schedule"
	"google.golang.org/grpc"
)

// NewScheduleServer initializes a new gRPC server for the Schedule service.
// It takes a gRPC server instance and a ScheduleRepository as parameters.
// It creates a new ScheduleService and a ScheduleHandler, and registers the handler with the gRPC server.
// This function is responsible for setting up the Schedule service in the gRPC server.
func NewScheduleServer(grpcServer *grpc.Server, repo *repository.ScheduleRepository) {
	service := &services.ScheduleService{Repo: *repo}
	handlers := &handler.ScheduleHandler{ScheduleService: service}
	schedule.RegisterScheduleServiceServer(grpcServer, handlers)
}
