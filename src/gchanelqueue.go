package gchanelqueue

import (
	"container/list"
	"fmt"
	"reflect"
)

// ChanelQueue by chanel
type ChanelQueue struct {
	sem  chan int
	list *list.List
}

func NewChanelQueue() *ChanelQueue {
	sem := make(chan int, 1)
	list := list.New()
	return &ChanelQueue{sem, list}
}

func (q *ChanelQueue) Size() int {
	return q.list.Len()
}

func (q *ChanelQueue) push(val interface{}) *list.Element {
	q.sem <- 1
	e := q.list.PushFront(val)
	<-q.sem
	return e
}

func (q *ChanelQueue) Pop() *list.Element {
	q.sem <- 1
	e := q.list.Back()
	q.list.Remove(e)
	<-q.sem
	return e
}
