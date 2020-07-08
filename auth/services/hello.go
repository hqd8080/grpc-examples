package services

import (
	"context"

	pb "github.com/hqd888/grpc-examples/auth/services/pb"
)

type HelloService struct{}

func (h *HelloService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	reply := &pb.HelloReply{Message: "hello " + request.GetName()}
	return reply, nil
}
