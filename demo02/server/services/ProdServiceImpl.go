package services

import (
	"context"
	"log"
)

type ProdService struct {
}

//服务具体实现
func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	log.Println(req.ProdId)
	return &ProdResponse{ProdStock: 20}, nil
}
