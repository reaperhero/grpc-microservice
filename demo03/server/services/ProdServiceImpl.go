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

func (this *ProdService) GetProdStocks(ctx context.Context, req *QuerySize) (*ProdResponseList, error) {
	log.Println(req.Size)
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 12},
		&ProdResponse{ProdStock: 13},
		&ProdResponse{ProdStock: 14},
	}
	return &ProdResponseList{Prodres: Prodres}, nil
}
