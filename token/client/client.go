package main

import (
	"context"
	"flag"
	"log"

	"github.com/hqd888/grpc-examples/token/services"
	pb "github.com/hqd888/grpc-examples/token/services/pb"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:8888", "addr")
	name = flag.String("name", "hqd888", "name")
)

type Authentication struct {
	Login    string
	Password string
}

func main() {
	flag.Parse()
	auth := services.Authentication{
		User:     "test",
		Password: "1234",
	}
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("failed to connect server,err:%v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Println(err)
	}

	log.Println(reply)
}
