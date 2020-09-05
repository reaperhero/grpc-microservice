package main

import (
	"google.golang.org/grpc"
	"grpc-microservice/demo03/server/services"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	err = rpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
