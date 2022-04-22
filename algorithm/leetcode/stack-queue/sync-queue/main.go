package sync_queue


type bbq struct {
	lok chan struct{}
	popLock chan struct{}
	data []interface{}
	len int
	size int
}

func NewBbq(size int) *bbq {
	bq := &bbq {
		lok : make(chan struct{}, 1),
		popLock: make(chan struct{}, 1),
		data : make([]interface{}, 0),
		len : 0,
		size: size,
	}
	bq.lok <- struct{}{}
	bq.popLock <- struct{}{}
	return bq
}

func (bq *bbq)Pop() interface{} {
    <- bq.lok
    if bq.len == 0{
    	return nil
	}
	res := bq.data[0]
	bq.data = bq.data[1:]
	bq.len--
	bq.lok <- struct{}{}
	if len(bq.popLock) == 0 {
		bq.popLock <- struct{}{}
	}
	return res
}

func (bq *bbq)Push(ele interface{}) {
	<- bq.popLock
	<- bq.lok
	bq.data = append(bq.data, ele)
	bq.len++
	if bq.len < bq.size {
		bq.popLock <- struct{}{}
	}
	bq.lok <- struct{}{}
}


// 未测试