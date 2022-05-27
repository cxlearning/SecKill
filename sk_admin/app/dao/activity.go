package dao

import "time"

type Activity struct {
	Dao
	ActivityId   string       `json:"activity_id" gorm:"type:varchar(128);not null;default:'';uniqueIndex:idx_aid;comment:活动ID"` //活动Id
	ActivityName string    `json:"activity_name" gorm:"type:varchar(128);not null;default:'';comment:活动名称"`                   //活动名称
	ProductId    string       `json:"product_id"  gorm:"type:varchar(128);not null;default:'';comment:商品Id"`                     //商品Id
	StartTime    time.Time `json:"start_time"`                                                                                //开始时间
	EndTime      time.Time `json:"end_time"`                                                                                  //结束时间
	Total        int       `json:"total"`                                                                                     //商品总数
	Status       int       `json:"status"`                                                                                    //状态

	Speed    int     `json:"speed"`
	BuyLimit int     `json:"buy_limit"`
	BuyRate  float64 `json:"buy_rate"`
}
