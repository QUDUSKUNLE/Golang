package handler

import (
	"github.com/QUDUSKUNLE/microservices/gateway/protogen/organization"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"google.golang.org/grpc"
)

type OrganizationServiceStruct struct {
	organizationService ports.UseCasePorts
	organization.UnimplementedOrganizationServiceServer
}

func NewOrganizationServer(grpcServer *grpc.Server, useCase ports.UseCasePorts) {
	organizationGrpc := &OrganizationServiceStruct{organizationService: useCase}
	organization.RegisterOrganizationServiceServer(grpcServer, organizationGrpc)
}
