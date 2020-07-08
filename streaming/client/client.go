package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/hqd888/grpc-examples/streaming/services/pb"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8888", "address")

func main() {
	flag.Parse()

	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect server,err:%v", err)
	}
	defer cc.Close()

	c := pb.NewUserServiceClient(cc)

	login(c)
	loginClientStreaming(c)
	loginServerStreaming(c)
	loginStreaming(c)
}

// 普通调用
func login(client pb.UserServiceClient) {
	req := &pb.UserLoginRequest{
		UserName: "hqd888",
		UserPwd:  "123456",
	}
	resp, err := client.Login(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

// 请求使用单向流
func loginClientStreaming(client pb.UserServiceClient) {
	login := &pb.UserLoginRequest{
		UserName: "hqd888",
		UserPwd:  "123456",
	}
	stream, err := client.LoginClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("failed to call,err:%v", err)
	}
	err = stream.Send(login)
	if err != nil {
		log.Fatalf("failed to send,err:%v", err)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to recv,err:%v", err)
	}
	log.Println(resp)
}

// 响应使用单向流
func loginServerStreaming(client pb.UserServiceClient) {
	login := &pb.UserLoginRequest{
		UserName: "hqd888",
		UserPwd:  "123",
	}
	stream, err := client.LoginServerStreaming(context.Background(), login)
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to recv,err:%v", err)
		}

		log.Println(resp)
	}
}

// 使用双向流
func loginStreaming(client pb.UserServiceClient) {
	req := &pb.UserLoginRequest{
		UserName: "hqd888",
		UserPwd:  "123456",
	}
	stream, err := client.LoginStreaming(context.Background())
	if err != nil {
		log.Fatalf("failed to call,err:%v", err)
	}

	err = stream.Send(req)
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
		log.Println(resp)
	}
}
