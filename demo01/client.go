package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-microservice/demo01/server/services"
	"log"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 10})
	if err != nil {
		log.Println(err)
	}
	log.Println(prodRes.ProdStock)
}
