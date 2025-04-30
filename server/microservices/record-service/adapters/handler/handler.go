package handler

import (
	organizationPorts "github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/record-service/clients"
	"github.com/QUDUSKUNLE/microservices/record-service/core/ports"
	v1 "github.com/QUDUSKUNLE/microservices/user-service/v1"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/record"
	"google.golang.org/grpc"
)

type RecordServiceStruct struct {
	recordService       ports.RecordPorts
	organizationService organizationPorts.OrganizationPorts
	userService         v1.UserPorts
	record.UnimplementedRecordServiceServer
}

func NewRecordServer(grpcServer *grpc.Server, useCase ports.RecordPorts, organi_conn, user_conn string) {
	// Use default client options (insecure) for internal service communication
	clientOptions := clients.DefaultClientOptions()

	recordGrpc := &RecordServiceStruct{
		recordService:       useCase,
		organizationService: clients.NewGRPClientOrganizationService(organi_conn, clientOptions),
		userService:         clients.NewGRPClientUserService(user_conn, clientOptions),
	}
	record.RegisterRecordServiceServer(grpcServer, recordGrpc)
}
