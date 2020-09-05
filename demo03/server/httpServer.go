package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpc-microservice/demo03/server/services"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	gwmux := runtime.NewServeMux()
	defer cancel()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := services.RegisterProdServiceHandlerFromEndpoint(ctx, gwmux, "localhost:8000", opt)
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":8001",
		Handler: gwmux,
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
