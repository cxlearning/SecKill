package model

type ProductCreatReq struct {
	ProductName string `json:"product_name" binding:"required"` //商品名称
	Total       int    `json:"total" binding:"required"`        //商品数量
	Status      int    `json:"status" binding:"required"`
}
