package clients

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/shared/db"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/organization"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	v1 "github.com/QUDUSKUNLE/microservices/user-service/v1"
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

// UpdateUser updates a user's information.
func (u *userService) UpdateUser(ctx context.Context, user db.UpdateUserParams) (*db.UpdateUserRow, error) {
	// Implement the logic to update a user using u.userGrpcClient
	panic("unimplemented")
}

func (u *userService) GetUsers(ctx context.Context, params db.GetUsersParams) ([]*db.User, error) {
	// Implement the logic to fetch users using u.userGrpcClient and params
	panic("unimplemented")
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	// Implement the logic to fetch user by email using u.userGrpcClient
	panic("unimplemented")
}

func NewGRPClientOrganizationService(organization_conn string, opts *ClientOptions) ports.OrganizationPorts {
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

func NewGRPClientUserService(user_conn string, opts *ClientOptions) v1.UserPorts {
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

	auth, err := grpc.Dial(user_conn, dialOpts...)
	if err != nil {
		log.Printf("Failed to connect to auth service: %v", err)
		return nil
	}
	return &userService{userGrpcClient: user.NewUserServiceClient(auth)}
}
