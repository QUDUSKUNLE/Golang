package client

import (
	"time"

	"github.com/QUDUSKUNLE/microservices/shared/protogen/organization"
)

const DefaultRequestTimeout = time.Second * 10

type organizationService struct {
	grpcClient organization.OrganizationServiceClient
}

// func NewGRPCOrganizationService(connString string) ports.OrganizationPorts {
// 	conn, err := grpc.NewClient(connString, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil
// 	}
// 	return &organizationService{grpcClient: organization.NewOrganizationServiceClient(conn)}
// }
