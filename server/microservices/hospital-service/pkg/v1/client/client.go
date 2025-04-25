package client

import (
	"time"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/organization"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const DefaultRequestTimeout = time.Second * 10

type organizationService struct {
	grpcClient organization.OrganizationServiceClient
}

func NewGRPCOrganizationService(connString string) ports.UseCasePorts {
	conn, err := grpc.NewClient(connString, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}
	return &organizationService{grpcClient: organization.NewOrganizationServiceClient(conn)}
}
