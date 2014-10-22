package base

import (
	//"container/list"
	"sync"
)

type CountDownLatch struct {
	cond  *sync.Cond
	count int
}

func NewCondDownLatch(count int) *CountDownLatch {
	lock := new(sync.Mutex)
	newcond := sync.NewCond(lock)
	return &CountDownLatch{newcond, count}
}

func (cdl *CountDownLatch) wait() {
	cdl.cond.L.Lock()
	defer cdl.cond.L.Unlock()
	for cdl.count > 0 {
		cdl.cond.Wait()
	}

}

func (cdl *CountDownLatch) countDown() {
	cdl.cond.L.Lock()
	defer cdl.cond.L.Unlock()
	cdl.count--
	if cdl.count == 0 {
		cdl.cond.Broadcast()
	}
}

func (cdl *CountDownLatch) getCount() int {
	cdl.cond.L.Lock()
	defer cdl.cond.L.Unlock()
	return cdl.count
}

//// GLockQueue by chanel
//type GCondBlockQueue struct {
//countdownlatch *CountDownLatch
//list  *list.List
//}

//func NewGLockQueue() *GCondBlockQueue {
//list := list.New()
//return &GLockQueue{NewCondDownLatch, list}
//}

//func (q *GLockQueue) Size() int {
//q.mutex.Lock()
//defer q.mutex.Unlock()
//return q.list.Len()
//}

//func (q *GLockQueue) Push(val interface{}) *list.Element {
//q.mutex.Lock()
//defer q.mutex.Unlock()
//e := q.list.PushFront(val)
//return e
//}

//func (q *GLockQueue) Pop() *list.Element {
//q.mutex.Lock()
//defer q.mutex.Unlock()
//e := q.list.Back()
//q.list.Remove(e)
//return e
/*}*/
