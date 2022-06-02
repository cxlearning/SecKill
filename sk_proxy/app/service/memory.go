package service

import (
	"fmt"
	"myProject/SecKill/sk_proxy/conf"
	"myProject/SecKill/sk_proxy/memory"
)

func Init() {
	memoryInit()
	processInit()
}

/**
初始化内存数据，加载商品信息
*/
func memoryInit() {

	memory.Mem.SecReqChan = make(chan *memory.SecRequest, 1000)
	memory.Mem.UserConns.UserConnMap = make(map[string]chan *memory.SecResponse, 1000)
	err := loadProductListFromEtcd(conf.Config.Etcd.EtcdSecProductKey)
	if err != nil {
		panic(fmt.Sprintf("load product list from etcd, err =%s", err.Error()))
	}
}

/**
处理线程初始化
1 将通道中req 放到redis队列中
2 将redis队列中结果 放到req的通道
*/
func processInit() {

	for i := 0; i < conf.Config.Server.WriteProxy2LayerGoroutineNum; i++{
		go WriteHandle()
	}

	for i := 0; i < conf.Config.Server.ReadProxy2LayerGoroutineNum; i++ {
		go ReadHandle()
	}


}
