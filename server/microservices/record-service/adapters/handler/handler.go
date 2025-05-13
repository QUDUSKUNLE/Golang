package handler

import (
	"github.com/QUDUSKUNLE/microservices/record-service/clients"
	"github.com/QUDUSKUNLE/microservices/record-service/core/ports"
	v1 "github.com/QUDUSKUNLE/microservices/user-service/adapters"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/record"
	"google.golang.org/grpc"
)

type RecordServiceStruct struct {
	recordService       ports.RecordPorts
	userService         v1.UserPorts
	fileService         LocalFileService
	record.UnimplementedRecordServiceServer
}

func NewRecordServer(grpcServer *grpc.Server, useCase ports.RecordPorts, organi_conn, user_conn string) {
	// Use default client options (insecure) for internal service communication
	clientOptions := clients.DefaultClientOptions()

	recordGrpc := &RecordServiceStruct{
		recordService:       useCase,
		userService:         clients.NewGRPClientUserService(user_conn, clientOptions),
	}
	record.RegisterRecordServiceServer(grpcServer, recordGrpc)
}
