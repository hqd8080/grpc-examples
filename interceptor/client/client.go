package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/hqd888/grpc-examples/interceptor/services/pb"
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

	for i := 1; i <= 10; i++ {
		reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("message:%d\t%s", i, reply.Message)
	}

	stream, err := client.Streaming(context.Background())
	if err != nil {
		log.Fatalf("failed to call,err:%v", err)
	}

	err = stream.Send(&pb.HelloRequest{Name: "streaming"})
	if err != nil {
		log.Fatalf("failed to send,err:%v", err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to recv,err:%v", err)
		}
		log.Printf("message:%s", resp.Message)
	}

}
