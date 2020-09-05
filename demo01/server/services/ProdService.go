package services

import (
	"context"
	"log"
)

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	// 打印客户端prod_id
	log.Println(in.ProdId)
	return &ProdResponse{ProdStock: 20}, nil
}
