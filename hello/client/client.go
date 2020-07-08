package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/hqd888/grpc-examples/hello/services/pb"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:8888", "addr")
	name = flag.String("name", "hqd888", "name")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect server,err:%v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)

}
