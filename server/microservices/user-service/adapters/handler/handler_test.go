package handler

import (
	"testing"

	"github.com/QUDUSKUNLE/microservices/events-service/domain"
	v1 "github.com/QUDUSKUNLE/microservices/user-service/adapters"
	"google.golang.org/grpc"
)

func TestNewAuthServer(t *testing.T) {
	type args struct {
		server  *grpc.Server
		usecase v1.UserPorts
		event   domain.EventPorts
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		NewUserService(tt.args.server, tt.args.usecase, tt.args.event)
	}
}
