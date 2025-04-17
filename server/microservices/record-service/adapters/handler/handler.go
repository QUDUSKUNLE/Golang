package handler

import (
	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	organizationPorts "github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/record-service/clients"
	"github.com/QUDUSKUNLE/microservices/record-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/record-service/protogen/golang/record"
	"google.golang.org/grpc"
)

type RecordServiceStruct struct {
	recordService       ports.UseCasePorts
	organizationService organizationPorts.UseCasePorts
	authService         v1.UseCaseInterface
	record.UnimplementedRecordServiceServer
}

func NewRecordServer(grpcServer *grpc.Server, useCase ports.UseCasePorts, organi_conn, auth_conn string) {
	// Use default client options (insecure) for internal service communication
	clientOptions := clients.DefaultClientOptions()
	
	recordGrpc := &RecordServiceStruct{
		recordService:       useCase,
		organizationService: clients.NewGRPClientOrganizationService(organi_conn, clientOptions),
		authService:         clients.NewGRPClientAuthService(auth_conn, clientOptions),
	}
	record.RegisterRecordServiceServer(grpcServer, recordGrpc)
}
