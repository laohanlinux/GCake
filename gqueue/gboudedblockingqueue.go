package gqueue

import (
	"container/list"
	"fmt"
	"sync"
)

type GBoudedBlockingQueue struct {
	mutex         *sync.Mutex // productor and customer all use the same Locker
	notEmpty      *Condition
	notFull       *Condition
	queue         *list.List
	queueCapacity int
}

func NewGBoundedBlockingQueue(size int) *GBoudedBlockingQueue {
	mutex := new(sync.Mutex)
	notEmpty_ := NewCondition(mutex)
	notFull_ := NewCondition(mutex)
	queue_ := new(list.List)
	return &GBoudedBlockingQueue{mutex, notEmpty_, notFull_, queue_, size}
}

func (gbq *GBoudedBlockingQueue) size() int {
	gbq.mutex.Lock()
	defer gbq.mutex.Unlock()
	return gbq.queue.Len()
}

func (gbq *GBoudedBlockingQueue) capacity() int {
	gbq.mutex.Lock()
	defer gbq.mutex.Unlock()
	return gbq.queueCapacity
}

func (gbq *GBoudedBlockingQueue) Put(value interface{}) {
	gbq.mutex.Lock()
	fmt.Println("Put Element: ", value)
	defer gbq.mutex.Unlock()
	for gbq.queue.Len() >= gbq.queueCapacity {
		fmt.Println("queue is full, ", value)
		gbq.notFull.wait()
	}
	gbq.queue.PushFront(value)
	gbq.notEmpty.notify()

}

func (gbq *GBoudedBlockingQueue) take() *list.Element {
	gbq.mutex.Lock()
	fmt.Println("Take a Element")
	defer gbq.mutex.Unlock()
	for gbq.queue.Len() == 0 {
		fmt.Println("queue is empty!!!")
		gbq.notEmpty.wait()
	}
	e := gbq.queue.Back()
	gbq.queue.Remove(e)
	gbq.notFull.notify()
	return e
}

func (gbq *GBoudedBlockingQueue) full() bool {
	gbq.mutex.Lock()
	defer gbq.mutex.Unlock()
	if gbq.queue.Len() > 10 {
		return true
	} else {
		return false
	}
}

func (gbq *GBoudedBlockingQueue) empty() bool {
	gbq.mutex.Lock()
	defer gbq.mutex.Unlock()
	if gbq.queue.Len() == 0 {
		fmt.Println("queue is empty")
		return true
	} else {
		fmt.Println("queue is not empty")
		return false
	}
}
