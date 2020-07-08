package services

import (
	"context"
	"io"

	pb "github.com/hqd888/grpc-examples/interceptor/services/pb"
)

type HelloService struct{}

func (h *HelloService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	reply := &pb.HelloReply{Message: "hello " + request.GetName()}
	return reply, nil
}

func (h *HelloService) Streaming(stream pb.HelloService_StreamingServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &pb.HelloReply{Message: "hello " + req.GetName()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
		return nil
	}
}
