package memory

import "sync"

/**
商品卖出数量统计，由于采用改数据结构所以layer层只能有一个
 */
type ProductCountMgr struct {
	productCount map[string]int
	lock         sync.RWMutex
}

func NewProductCountMgr() *ProductCountMgr {
	productMgr := &ProductCountMgr{
		productCount: make(map[string]int, 128),
	}
	return productMgr
}

/**
获得商品已售数量
 */
func (p *ProductCountMgr) Count(productId string) (count int) {
	p.lock.RLock()
	defer p.lock.Unlock()

	count = p.productCount[productId]
	return
}

//添加商品
func (p *ProductCountMgr) Add(productId string, count int) {
	p.lock.Lock()
	defer p.lock.Unlock()

	cur, ok := p.productCount[productId]
	if !ok {
		cur = count
	} else {
		cur += count
	}
	p.productCount[productId] = cur
}