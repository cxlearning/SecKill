package service

import (
	"github.com/satori/go.uuid"
	"myProject/SecKill/sk_admin/app/dao"
	"myProject/SecKill/sk_admin/app/model"
)

var ProductService productService

type productService struct {
}

func (p *productService) Creat(req model.ProductCreatReq) error {

	product := dao.Product{
		ProductId:   uuid.NewV4().String(),
		ProductName: req.ProductName,
		Total:       req.Total,
		Status:      req.Status,
	}

	if err := product.Save(); err != nil {
		return err
	}
	return nil
}
