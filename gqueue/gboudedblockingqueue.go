package gqueue

import (
	"container/list"
	"sync"
)

type GBoudedBlockingQueue struct {
	mutex    *sync.Mutex
	notEmpty *Condition
	notFull  *Condition
	queue    *list.List
}

func NewGBoundedBlockingQueue(mutex ...sync.Mutex) *GBoudedBlockingQueue {
	var mutex_ *sync.Mutex
	if len(mutex) == 0 {
		mutex_ = new(sync.Mutex)
	} else {
		*mutex_ = mutex[0]
	}
	notEmpty_ := NewCondition(mutex_)
	notFull_ := NewCondition(mutex_)
	queue_ := new(list.List)
	return &GBoudedBlockingQueue{mutex_, notEmpty_, notFull_, queue_}
}

func (gbq *GBoudedBlockingQueue) size() int {
	gbq.mutex.Lock()
	defer gbq.mutex.Unlock()
	return gbq.queue.Len()
}
