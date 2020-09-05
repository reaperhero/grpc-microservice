package main

import (
	"google.golang.org/grpc"
	"grpc-microservice/demo01/server/services"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, _ := net.Listen("tcp", ":8000")
	rpcServer.Serve(lis)
}