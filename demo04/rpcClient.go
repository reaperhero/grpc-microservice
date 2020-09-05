package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-microservice/demo04/server/services"
	"log"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	prodresA, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 10, ProdArea: services.ProdAreas_A})
	prodresB, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 10, ProdArea: services.ProdAreas_B})
	prodresC, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 10, ProdArea: services.ProdAreas_C})
	if err != nil {
		log.Println(err)
	}
	log.Println(prodresA.ProdStock)
	log.Println(prodresB.ProdStock)
	log.Println(prodresC.ProdStock)
}
