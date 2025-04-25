package handler

import (
	"testing"

	v1 "github.com/QUDUSKUNLE/microservices/auth-service/pkg/v1"
	"google.golang.org/grpc"
)

func TestNewAuthServer(t *testing.T) {
	type args struct {
		server  *grpc.Server
		usecase v1.UseCaseInterface
		conn    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewAuthServer(tt.args.server, tt.args.usecase, tt.args.conn)
		})
	}
}
