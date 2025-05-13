package handler

import (
	"context"
	"reflect"
	"testing"

	v1 "github.com/QUDUSKUNLE/microservices/auth-service/adapters"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/user"
	"github.com/QUDUSKUNLE/microservices/shared/protogen/auth"
)

type fields struct {
	authService                    v1.AuthPorts
	UnimplementedUserServiceServer user.UnimplementedUserServiceServer
}

func TestUserServiceStruct_Signin(t *testing.T) {
	type args struct {
		ctx context.Context
		req *auth.SignInRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *auth.SignInResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &AuthServiceStruct{
				authService:                    tt.fields.authService,
				// Removed as AuthServiceStruct does not have this field
			}
			got, err := srv.Signin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthServiceStruct.Signin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthServiceStruct.Signin() = %v, want %v", got, tt.want)
			}
		})
	}
}
