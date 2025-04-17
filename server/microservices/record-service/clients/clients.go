package clients

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/QUDUSKUNLE/microservices/auth-service/adapters/db"
	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/auth-service/protogen/golang/user"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/organization-service/protogen/golang/organization"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const DefaultRequestTimeout = time.Second * 10

type ClientOptions struct {
	UseTLS       bool
	CertFilePath string
}

func DefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		UseTLS:       false,
		CertFilePath: "",
	}
}

func WithTLS(certFilePath string) *ClientOptions {
	return &ClientOptions{
		UseTLS:       true,
		CertFilePath: certFilePath,
	}
}

type organizationService struct {
	organizationGrpcClient organization.OrganizationServiceClient
}

type userService struct {
	userGrpcClient user.UserServiceClient
}

func (u *userService) GetUsers(ctx context.Context) ([]*db.User, error) {
	panic("unimplemented")
}

func NewGRPClientOrganizationService(organization_conn string, opts *ClientOptions) ports.UseCasePorts {
	if opts == nil {
		opts = DefaultClientOptions()
	}

	var dialOpts []grpc.DialOption

	if opts.UseTLS {
		config := &tls.Config{
			ServerName:         "localhost",
			InsecureSkipVerify: true, // For development only
			MinVersion:         tls.VersionTLS12,
		}
		creds := credentials.NewTLS(config)
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	organi_conn, err := grpc.Dial(organization_conn, dialOpts...)
	if err != nil {
		log.Printf("Failed to connect to organization service: %v", err)
		return nil
	}
	return &organizationService{organizationGrpcClient: organization.NewOrganizationServiceClient(organi_conn)}
}

func NewGRPClientAuthService(auth_conn string, opts *ClientOptions) v1.UseCaseInterface {
	if opts == nil {
		opts = DefaultClientOptions()
	}

	var dialOpts []grpc.DialOption

	if opts.UseTLS {
		config := &tls.Config{
			ServerName:         "localhost",
			InsecureSkipVerify: true, // For development only
			MinVersion:         tls.VersionTLS12,
		}
		creds := credentials.NewTLS(config)
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	} else {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	auth, err := grpc.Dial(auth_conn, dialOpts...)
	if err != nil {
		log.Printf("Failed to connect to auth service: %v", err)
		return nil
	}
	return &userService{userGrpcClient: user.NewUserServiceClient(auth)}
}
