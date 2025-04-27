package handler

import (
	"reflect"
	"testing"

	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
)

func TestUserServiceStruct_transformUserRPC(t *testing.T) {
	type fields struct {
		authService                    v1.AuthPorts
		organizationService            ports.OrganizationPorts
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
	}
	type args struct {
		req *user.CreateUserRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   dto.UserDto
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &AuthServiceStruct{
				authService:                    tt.fields.authService,
				organizationService:            tt.fields.organizationService,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			if got := srv.transformUserRPC(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthServiceStruct.transformUserRPC() = %v, want %v", got, tt.want)
			}
		})
	}
}
