package services

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hqd888/grpc-examples/token/services/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type Authentication struct {
	User     string
	Password string
}

type HelloService struct {
	auth *Authentication
}

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

// 是否使用安全链接
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

// 身份认证
func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string

	log.Println(md)

	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	log.Println(appid)
	log.Println(appkey)

	if appid != "test" || appkey != "1234" {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil

}

func (h *HelloService) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	if err := h.auth.Auth(ctx); err != nil {
		return nil, err
	}

	reply := &pb.HelloReply{Message: "hello " + request.GetName()}
	return reply, nil
}
