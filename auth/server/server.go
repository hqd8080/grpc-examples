package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/hqd888/grpc-examples/auth/services"
	pb "github.com/hqd888/grpc-examples/auth/services/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr = flag.String("addr", ":8888", "address")

func main() {
	flag.Parse()

	certificate, err := tls.LoadX509KeyPair("../keys/server.crt", "../keys/server.key")
	if err != nil {
		log.Fatalf("failed to load files,err:%v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../keys/ca.crt")
	if err != nil {
		log.Fatalf("failed to read file,err:%v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		//ServerName:   "tlsServerName",
		ClientCAs: certPool,
	})

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen,err:%v", err)
	}

	server := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServiceServer(server, new(services.HelloService))

	log.Printf("server start at [%s]", *addr)

	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
