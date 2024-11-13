package grpc

import (
	"context"

	pb "github.com/QUDUSKUNLE/microservices/services/auth-service/protogen/golang/greet"
)

func (srv *UserServiceStruct) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return srv.transformMessage(req), nil
}
