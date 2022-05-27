package model

type ActivityCreatReq struct {
	ActivityName string  `json:"activity_name" gorm:"type:varchar(128);not null;default:'';comment:活动名称"` //活动名称
	ProductId    string  `json:"product_id"  gorm:"type:varchar(128);not null;default:'';comment:商品Id"`   //商品Id
	StartTime    int64   `json:"start_time"`                                                              //开始时间
	EndTime      int64   `json:"end_time"`                                                                //结束时间
	Total        int     `json:"total"`                                                                   //商品总数
	Status       int     `json:"status"`                                                                  //状态
	Speed        int     `json:"speed"`
	BuyLimit     int     `json:"buy_limit"`
	BuyRate      float64 `json:"buy_rate"`
}

/**
etcd 中存储的数据结构
*/
type SecProductInfoConf struct {
	ProductId         string  `json:"product_id"`           //商品Id
	StartTime         int64   `json:"start_time"`           //开始时间
	EndTime           int64   `json:"end_time"`             //结束时间
	Status            int     `json:"status"`               //状态
	Total             int     `json:"total"`                //商品总数
	Left              int     `json:"left"`                 //剩余商品数
	OnePersonBuyLimit int     `json:"one_person_buy_limit"` //一个人购买限制
	BuyRate           float64 `json:"buy_rate"`             //买中几率
	SoldMaxLimit      int     `json:"sold_max_limit"`       //每秒最多能卖多少个
}
