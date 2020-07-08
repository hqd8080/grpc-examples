package main

import (
	"flag"
	"log"
	"net"

	"github.com/hqd888/grpc-examples/token/services"
	pb "github.com/hqd888/grpc-examples/token/services/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr = flag.String("addr", ":8888", "address")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen,err:%v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterHelloServiceServer(srv, new(services.HelloService))
	reflection.Register(srv)

	log.Printf("server start at [%s]", *addr)

	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
