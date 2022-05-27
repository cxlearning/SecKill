package dao

import (
	"myProject/SecKill/sk_admin/library/db"
)

type Activity struct {
	Dao
	ActivityId   string  `json:"activity_id" gorm:"type:varchar(128);not null;default:'';uniqueIndex:idx_aid;comment:活动ID"` //活动Id
	ActivityName string  `json:"activity_name" gorm:"type:varchar(128);not null;default:'';comment:活动名称"`                   //活动名称
	ProductId    string  `json:"product_id"  gorm:"type:varchar(128);not null;default:'';comment:商品Id"`                     //商品Id
	StartTime    int64   `json:"start_time"`                                                                                //开始时间
	EndTime      int64   `json:"end_time"`                                                                                  //结束时间
	Total        int     `json:"total"`                                                                                     //商品总数
	Status       int     `json:"status"`                                                                                    //状态
	Speed        int     `json:"speed"`                                                                                     // 每秒可以卖几个
	BuyLimit     int     `json:"buy_limit"`                                                                                 // 每人的购买限制
	BuyRate      float64 `json:"buy_rate"`                                                                                  // 买中几率
}

func (a *Activity) Save() error {
	return db.GetDBInstance().Save(a).Error
}
