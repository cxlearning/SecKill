package memory

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"log"
	"myProject/SecKill/sk_layer/conf"
	"myProject/SecKill/sk_layer/library/etcd"
)


/**
etcd 对内存的操作
 */

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

	go watchSecProductKey()
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

func watchSecProductKey() {
	key := conf.Config.Etcd.EtcdSecProductKey

	watchch := etcd.GetEtcdInstance().Watch(context.Background(), key)

	for _w := range watchch {
		flag := false
		var secProductInfoList []*SecProductInfoConf
		for _, _e := range _w.Events { // 这里都是同一个key的事件
			//删除事件
			if _e.Type == mvccpb.DELETE {
				log.Printf("key[%s] 's config deleted", key)
				continue
			}
			if _e.Type == mvccpb.PUT && string(_e.Kv.Key) == key{
				err := json.Unmarshal(_e.Kv.Value, &secProductInfoList)
				if err != nil {
					log.Printf("key [%s], Unmarshal[%s]. Error : %v", key, err)
					continue
				}
				flag = true
			}

			if flag {
				log.Printf("get config from etcd success, %v", secProductInfoList)
				updatProducts(secProductInfoList)
			}

			log.Printf("get config from etcd, %s %q : %q\n", _e.Type, _e.Kv.Key, _e.Kv.Value)
		}
	}
}
