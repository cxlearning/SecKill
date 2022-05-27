package service

import (
	"context"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"log"
	"myProject/SecKill/sk_admin/app/dao"
	"myProject/SecKill/sk_admin/app/model"
	"myProject/SecKill/sk_admin/conf"
	"myProject/SecKill/sk_admin/library/etcd"
)

var ActivityService activityService

type activityService struct {
}

func (a *activityService) Creat(req model.ActivityCreatReq) error {

	activity := dao.Activity{
		Dao:          dao.Dao{},
		ActivityId:   uuid.NewV4().String(),
		ActivityName: req.ActivityName,
		ProductId:    req.ProductId,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Total:        req.Total,
		Status:       req.Status,
		Speed:        req.Status,
		BuyLimit:     req.BuyLimit,
		BuyRate:      req.BuyRate,
	}

	if err := activity.Save(); err != nil {
		return err
	}
	err := a.syncToEtcd(activity)
	if err != nil {
		return err
	}

	return nil
}

func (a *activityService) syncToEtcd(activity dao.Activity) error {

	productInfoList, err := a.loadProductListFromEtcd(conf.Config.Etcd.EtcdSecActivityKey)
	if err != nil {
		return err
	}
	proInfo := model.SecProductInfoConf{
		ProductId:         activity.ProductId,
		StartTime:         activity.StartTime,
		EndTime:           activity.EndTime,
		Status:            activity.Status,
		Total:             activity.Total,
		Left:              activity.Total,
		OnePersonBuyLimit: activity.BuyLimit,
		BuyRate:           activity.BuyRate,
		SoldMaxLimit:      activity.Speed,
	}
	productInfoList = append(productInfoList, &proInfo)

	data, err := json.Marshal(productInfoList)
	if err != nil {
		log.Printf("json marshal failed, err : %v", err)
		return err
	}

	_, err = etcd.GetEtcdInstance().Put(context.Background(), conf.Config.Etcd.EtcdSecActivityKey, string(data))
	if err != nil {
		log.Printf("put to etcd failed, err : %v, data = [%v]", err, string(data))
		return err
	}

	log.Printf("put to etcd success, data = [%v]", string(data))
	return nil

}

//从Ectd中取出原来的商品数据
func (a *activityService) loadProductListFromEtcd(key string) ([]*model.SecProductInfoConf, error) {

	log.Println("start get from etcd")

	resp, err := etcd.GetEtcdInstance().Get(context.Background(), key)
	if err != nil {
		log.Printf("get [%s] from etcd failed, err : %v", key, err)
		return nil, err
	}

	var secProductInfo []*model.SecProductInfoConf
	for k, v := range resp.Kvs {

		log.Printf("key = [%v], value = [%v]", k, v)
		err := json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			log.Printf("Unmsharl second product info failed, err : %v", err)
			return nil, err
		}
		log.Printf("second info conf is [%v]", secProductInfo)

	}
	return secProductInfo, nil
}
