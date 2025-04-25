package handler

import (
	"reflect"
	"testing"

	"github.com/QUDUSKUNLE/microservices/auth-service/adapters/dto"
	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"github.com/QUDUSKUNLE/microservices/gateway/protogen/user"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
)

func TestUserServiceStruct_transformUserRPC(t *testing.T) {
	type fields struct {
		userService                    v1.UseCaseInterface
		organizationService            ports.UseCasePorts
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
			srv := &UserServiceStruct{
				userService:                    tt.fields.userService,
				organizationService:            tt.fields.organizationService,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			if got := srv.transformUserRPC(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServiceStruct.transformUserRPC() = %v, want %v", got, tt.want)
			}
		})
	}
}
