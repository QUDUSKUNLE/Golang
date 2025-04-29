package handler

import (
	"testing"

	v1 "github.com/QUDUSKUNLE/microservices/user-service/pkg/v1"
	"google.golang.org/grpc"
)

func TestNewAuthServer(t *testing.T) {
	type args struct {
		server  *grpc.Server
		usecase v1.UserPorts
		// event   domain.EventPorts
		conn string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		NewAuthServer(tt.args.server, tt.args.usecase, tt.args.conn)
	}
}
