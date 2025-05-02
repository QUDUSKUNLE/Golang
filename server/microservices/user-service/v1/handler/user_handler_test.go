package handler

import (
	"context"
	"reflect"
	"testing"

	v1 "github.com/QUDUSKUNLE/microservices/user-service/v1"
	"github.com/QUDUSKUNLE/microservices/organization-service/core/ports"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
)

type fields struct {
	userService                    v1.UserPorts
	organizationService            ports.OrganizationPorts
	UnimplementedUserServiceServer user.UnimplementedUserServiceServer
}

func TestUserServiceStruct_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		req *user.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.SuccessResponse
		wantErr bool
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
			got, err := srv.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServiceStruct.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServiceStruct.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserServiceStruct_Read(t *testing.T) {
	type fields struct {
		userService                    v1.UserPorts
		organizationService            ports.OrganizationPorts
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *user.SingleUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.GetUserResponse
		wantErr bool
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
			got, err := srv.Read(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServiceStruct.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServiceStruct.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserServiceStruct_ReadUsers(t *testing.T) {
	type fields struct {
		userService                    v1.UserPorts
		organizationService            ports.OrganizationPorts
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *user.GetUsersRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.GetUsersResponse
		wantErr bool
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
			got, err := srv.ReadUsers(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServiceStruct.ReadUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServiceStruct.ReadUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserServiceStruct_UpdateNin(t *testing.T) {
	type fields struct {
		userService                    v1.UserPorts
		organizationService            ports.OrganizationPorts
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *user.UpdateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.UpdateUserResponse
		wantErr bool
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
			got, err := srv.UpdateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServiceStruct.UpdateNin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServiceStruct.UpdateNin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserServiceStruct_Home(t *testing.T) {
	type fields struct {
		userService                    v1.UserPorts
		organizationService            ports.OrganizationPorts
		UnimplementedUserServiceServer user.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *user.HomeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.GetHomeResponse
		wantErr bool
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
			got, err := srv.Home(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserServiceStruct.Home() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserServiceStruct.Home() = %v, want %v", got, tt.want)
			}
		})
	}
}
