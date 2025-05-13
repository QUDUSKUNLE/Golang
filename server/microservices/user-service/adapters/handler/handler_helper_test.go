package handler

import (
	"reflect"
	"testing"

	v1 "github.com/QUDUSKUNLE/microservices/user-service/adapters"
	"github.com/QUDUSKUNLE/microservices/shared/dto"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
)

func TestUserServiceStruct_transformUserRPC(t *testing.T) {
	type fields struct {
		userService                    v1.UserPorts
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
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			if got := srv.transformUserRPC(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServiceStruct.transformUserRPC() = %v, want %v", got, tt.want)
			}
		})
	}
}
