package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-microservice/demo03/server/services"
	"log"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStocks(context.Background(), &services.QuerySize{Size: 4})
	if err != nil {
		log.Println(err)
	}
	log.Println(prodRes.Prodres)
}
