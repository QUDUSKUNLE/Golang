package handler

import (
	"github.com/QUDUSKUNLE/microservices/shared/protogen/organization"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"google.golang.org/grpc"
)

type OrganizationServiceStruct struct {
	organizationService ports.OrganizationPorts
	organization.UnimplementedOrganizationServiceServer
}

func NewOrganizationServer(grpcServer *grpc.Server, useCase ports.OrganizationPorts) {
	organizationGrpc := &OrganizationServiceStruct{organizationService: useCase}
	organization.RegisterOrganizationServiceServer(grpcServer, organizationGrpc)
}
