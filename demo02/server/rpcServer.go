package main

import (
	"google.golang.org/grpc"
	"grpc-microservice/demo02/server/services"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, err := net.Listen("tcp",":8000")
	if err != nil{
		log.Fatal(err)
	}
	rpcServer.Serve(lis)
}
