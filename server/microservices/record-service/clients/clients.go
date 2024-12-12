package clients

import (
	"time"

	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const DefaultRequestTimeout = time.Second * 10

type organizationService struct {
	organizationGrpcClient organization.OrganizationServiceClient
}

type userService struct {
	userGrpcClient user.UserServiceClient
}

func NewGRPClientOrganizationService(organization_conn string) ports.UseCasePorts {
	organi_conn, err := grpc.NewClient(organization_conn, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}
	return &organizationService{organizationGrpcClient: organization.NewOrganizationServiceClient(organi_conn)}
}

func NewGRPClientAuthService(auth_conn string) v1.UseCaseInterface {
	auth, err := grpc.NewClient(auth_conn, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}
	return &userService{userGrpcClient: user.NewUserServiceClient(auth)}
}
