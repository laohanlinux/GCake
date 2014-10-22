package base

import (
	"container/list"
	//"reflect"
)

// ChanelQueue by chanel
type GChanelQueue struct {
	sem  chan int
	list *list.List
}

func NewChanelQueue() *GChanelQueue {
	sem := make(chan int, 1)
	list := list.New()
	return &GChanelQueue{sem, list}
}

func (q *GChanelQueue) Size() int {
	return q.list.Len()
}

func (q *GChanelQueue) Push(val interface{}) *list.Element {
	q.sem <- 1
	e := q.list.PushFront(val)
	<-q.sem
	return e
}

func (q *GChanelQueue) Pop() *list.Element {
	q.sem <- 1
	e := q.list.Front()
	q.list.Remove(e)
	<-q.sem
	return e
}
