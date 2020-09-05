package services

import (
	"context"
)

type ProdService struct {
}

//服务具体实现
func (this *ProdService) GetProdInfo(ctx context.Context, req *ProdRequest) (*ProdModel, error) {
	return &ProdModel{ProdId: 101, ProdName: "testname", ProdPrice: 20.5}, nil
}
