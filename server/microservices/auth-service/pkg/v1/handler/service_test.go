package handler

import (
	"testing"

	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"google.golang.org/grpc"
)

func TestNewAuthServer(t *testing.T) {
	type args struct {
		server   *grpc.Server
		authcase v1.AuthPorts
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		NewAuthServer(tt.args.server, tt.args.authcase)
	}
}
