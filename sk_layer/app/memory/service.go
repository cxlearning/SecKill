package memory

import (
	"context"
	"encoding/json"
	"log"
	"myProject/SecKill/sk_layer/conf"
	"myProject/SecKill/sk_layer/library/etcd"
)

func Init() {
	loadProductListFromEtcd(conf.Config.Etcd.EtcdSecProductKey)
}

//从Ectd中取出原来的商品数据加载到内存中，并热更新
func loadProductListFromEtcd(key string) error {

	log.Println("start get from etcd")

	resp, err := etcd.GetEtcdInstance().Get(context.Background(), key)
	if err != nil {
		log.Printf("get [%s] from etcd failed, err : %v", key, err)
		return err
	}

	var secProductInfoList []*SecProductInfoConf
	for k, v := range resp.Kvs {

		log.Printf("key = [%v], value = [%v]", k, v)
		err := json.Unmarshal(v.Value, &secProductInfoList)
		if err != nil {
			log.Printf("Unmsharl second product info failed, err : %v", err)
			return err
		}
		log.Printf("second info conf is [%v]", secProductInfoList)

	}

	// 更新内存
	updatProducts(secProductInfoList)

	// todo 热更新
	return nil
}

func updatProducts(products []*SecProductInfoConf) {

	m := make(map[string]*SecProductInfoConf, len(products))

	for _, _p := range products {
		m[_p.ProductId] = _p
	}

	Mem.Products.Lock.Lock()
	defer Mem.Products.Lock.Unlock()

	Mem.Products.ProductMap = m

}
