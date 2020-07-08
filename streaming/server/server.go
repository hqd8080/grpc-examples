package main

import (
	"flag"
	"log"
	"net"

	"github.com/hqd888/grpc-examples/streaming/services"
	pb "github.com/hqd888/grpc-examples/streaming/services/pb"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8888", "address")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen,err:%s", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, new(services.UserService))

	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
