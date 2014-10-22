package base

import (
	"container/list"
)

// GLockQueue by chanel
type GBlockingQueue struct {
	notEmpty *Condition
	list     *list.List
}

func NewGBlockingQueue() *GBlockingQueue {
	cond := NewCondition()
	list := list.New()
	return &GBlockingQueue{cond, list}
}

func (gbq *GBlockingQueue) Size() int {
	gbq.notEmpty.cond.L.Lock()
	defer gbq.notEmpty.cond.L.Unlock()
	return gbq.list.Len()
}

func (gbq GBlockingQueue) Put(val interface{}) {
	gbq.notEmpty.cond.L.Lock()
	defer gbq.notEmpty.cond.L.Unlock()
	gbq.list.PushFront(val)
	gbq.notEmpty.notify()
}

func (gbq GBlockingQueue) take() *list.Element {
	gbq.notEmpty.cond.L.Lock()
	defer gbq.notEmpty.cond.L.Unlock()
	for gbq.list.Len() == 0 {
		gbq.notEmpty.wait()
	}
	e := gbq.list.Back()
	gbq.list.Remove(e)
	return e
}
