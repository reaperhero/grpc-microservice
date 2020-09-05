package services

import (
	"context"
	"log"
)

type ProdService struct {
}

//服务具体实现
func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch req.ProdArea {
	case ProdAreas_B:
		stock = 31 //B区库存31个
	case ProdAreas_C:
		stock = 50 //C区库存50
	default:
		stock = 102 //默认A区，库存102
	}
	log.Println(req.ProdId, req.ProdArea)
	return &ProdResponse{ProdStock: stock}, nil
}
