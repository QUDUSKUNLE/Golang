package server

import (
	"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/handler"
		"github.com/QUDUSKUNLE/microservices/diagnostic-service/adapters/repository"
	"github.com/QUDUSKUNLE/microservices/diagnostic-service/core/services"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/diagnostic"
	"google.golang.org/grpc"
)

func NewDiagnosticServer(grpcServer *grpc.Server, repo *repository.DiagnosticRepository) {
	service := &services.DiagnosticService{Repo: *repo}
	handlers := &handler.DiagnosticHandler{DiagnosticService: service}
	diagnostic.RegisterDiagnosticServiceServer(grpcServer, handlers)
}
