package service

import (
	"fmt"
	"myProject/SecKill/sk_layer/conf"
	"myProject/SecKill/sk_layer/memory"
)

func Init()  {
	memoryInit()
	processInit()
}

/**
初始化内存数据，加载商品信息
 */
func memoryInit() {

	memory.Mem.Read2HandleChan = make(chan *memory.SecRequest, 1000)
	memory.Mem.Handle2WriteChan = make(chan *memory.SecResponse, 1000)
	err := loadProductListFromEtcd(conf.Config.Etcd.EtcdSecProductKey)
	if err != nil {
		panic(fmt.Sprintf("load product list from etcd, err =%s", err.Error()))
	}
}

/**
处理线程初始化
1 redis中读取请求放到通道中
2 处理通道中请求，并将响应放到通道中
3 将通道中响应放到redis中
 */
func processInit() {



}
