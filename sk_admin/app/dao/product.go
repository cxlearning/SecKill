package dao

type Product struct {
	Dao
	ProductId   string  `json:"product_id" gorm:"type:varchar(128);not null;default:'';uniqueIndex:idx_pid;comment:商品ID"`
	ProductName string `json:"product_name" gorm:"type:varchar(128);not null;default:'';comment:商品名称"` //商品名称
	Total       int    `json:"total" gorm:"type:int;not null;default:0;comment:商品数量"`        //商品数量
	Status      int    `json:"status" gorm:"type:int;not null;default:0;uniqueIndex:idx_pid;comment:商品状态"`       //商品状态
}
