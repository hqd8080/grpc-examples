#!/bin/bash
protoc -I=./proto --go_out=plugins=grpc:./services/pb  ./proto/*.proto