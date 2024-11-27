package handler

import (
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"google.golang.org/grpc"
)

type OrganizationServiceStruct struct {
	useCase ports.UseCasePorts
	organization.UnimplementedOrganizationServiceServer
}

func NewOrganizationServer(grpcServer *grpc.Server, useCase ports.UseCasePorts) {
	organizationGrpc := &OrganizationServiceStruct{useCase: useCase}
	organization.RegisterOrganizationServiceServer(grpcServer, organizationGrpc)
}
