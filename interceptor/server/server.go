package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/hqd888/grpc-examples/interceptor/services"
	pb "github.com/hqd888/grpc-examples/interceptor/services/pb"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8888", "address")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen,err:%v", err)
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryFilter),
		grpc.StreamInterceptor(StreamFilter),
	)

	pb.RegisterHelloServiceServer(srv, new(services.HelloService))

	log.Printf("server start at [%s]", *addr)

	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

func UnaryFilter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter:", info.FullMethod)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func StreamFilter(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("filter:%+v", info)

	err := handler(srv, ss)
	return err
}
