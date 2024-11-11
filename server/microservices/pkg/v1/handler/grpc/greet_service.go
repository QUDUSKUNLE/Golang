package grpc

import (
	"context"

	pb "github.com/QUDUSKUNLE/microservices/proto"
)

func (srv *UserServiceStruct) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return srv.transformMessage(req), nil
}
