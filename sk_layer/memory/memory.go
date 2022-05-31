package memory

import (
	"sync"
)

var Mem Memory

/**
存储在内存的数据结构
*/

type Memory struct {
	Products Products

	Read2HandleChan  chan *SecRequest
	Handle2WriteChan chan *SecResponse
}

//请求
type SecRequest struct {
	ProductId     int    `json:"product_id"` //商品ID
	Source        string `json:"source"`
	AuthCode      string `json:"auth_code"`
	SecTime       string `json:"sec_time"`
	Nance         string `json:"nance"`
	UserId        int    `json:"user_id"`
	UserAuthSign  string `json:"user_auth_sign"`
	AccessTime    int64  `json:"access_time"`
	ClientAddr    string `json:"client_addr"`
	ClientRefence string `json:"client_refence"`
}

//响应
type SecResponse struct {
	ProductId int    `json:"product_id"` //商品ID
	UserId    int    `json:"user_id"`    //用户ID
	Token     string `json:"token"`      //Token
	TokenTime int64  `json:"token_time"` //Token生成时间
	Code      int    `json:"code"`       //状态码
}

//秒杀活动信息，从etcd 中加载所得
type SecProductInfoConf struct {
	ProductId         string  `json:"product_id"`           //商品ID
	StartTime         int64   `json:"start_time"`           //秒杀开始时间
	EndTime           int64   `json:"end_time"`             //秒杀结束时间
	Status            int     `json:"status"`               //状态
	Total             int     `json:"total"`                //商品总数
	Left              int     `json:"left"`                 //商品剩余数量
	OnePersonBuyLimit int     `json:"one_person_buy_limit"` //单个用户购买数量限制
	BuyRate           float64 `json:"buy_rate"`             //买中几率
	SoldMaxLimit      int     `json:"sold_max_limit"`       //每秒最多能卖多少个
	//	SecLimit          *srv_limit.SecLimit `json:"sec_limit"` //限速控制
}

/**
所有的秒杀商品
*/
type Products struct {
	ProductMap map[string]*SecProductInfoConf // productID -- > SecProductInfoConf
	Lock       sync.RWMutex
}
