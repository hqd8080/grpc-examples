#!/bin/bash
#生成自签名证书
#openssl req -newkey rsa:2048 -nodes -keyout keys/server.key -x509 -days 3650 -out keys/server.crt -subj "/C=CN/L=China/O=grpc-server/CN=server.grpc.io"
#openssl req -newkey rsa:2048 -nodes -keyout keys/client.key -x509 -days 3650 -out keys/client.crt -subj "/C=CN/L=China/O=grpc-client/CN=client.grpc.io"

#根证书生成
#openssl genrsa -out keys/ca.key 2048
#openssl req -new -x509 -days 3650 -subj "/C=CN/L=China/O=grpc/CN=github.com" -key keys/ca.key -out keys/ca.crt

#服务器端证书签名
#openssl req -new -subj "/C=CN/L=China/O=grpc-server/CN=server.grpc.io" -key keys/server.key -out keys/server.csr
#openssl x509 -req -sha256 -CA keys/ca.crt -CAkey keys/ca.key -CAcreateserial -days 3650 -in keys/server.csr -out keys/server.crt

#客户端证书签名
#openssl req -new -subj "/C=CN/L=China/O=grpc-client/CN=client.grpc.io" -key keys/client.key -out keys/client.csr
#openssl x509 -req -sha256 -CA keys/ca.crt -CAkey keys/ca.key -CAcreateserial -days 3650 -in keys/client.csr -out keys/client.crt


protoc -I=./proto --go_out=plugins=grpc:./services/pb  ./proto/*.proto