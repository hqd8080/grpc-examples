package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"

	pb "github.com/hqd888/grpc-examples/auth/services/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addr = flag.String("addr", ":8888", "address")
	name = flag.String("name", "hqd888", "name")
)

func main() {
	flag.Parse()

	certificate, err := tls.LoadX509KeyPair("../keys/client.crt", "../keys/client.key")
	if err != nil {
		log.Fatalf("failed to load files,err:%v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../keys/ca.crt")
	if err != nil {
		log.Fatalf("failed to read file,err:%v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		//ServerName:   "tlsServerName",
		RootCAs: certPool,
	})

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
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
