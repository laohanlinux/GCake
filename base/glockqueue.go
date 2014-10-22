package base

import (
	"container/list"
	"sync"
)

// GLockQueue by chanel
type GLockQueue struct {
	mutex *sync.Mutex
	list  *list.List
}

func NewGLockQueue() *GLockQueue {
	lock := new(sync.Mutex)
	list := list.New()
	return &GLockQueue{lock, list}
}

func (q *GLockQueue) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.list.Len()
}

func (q *GLockQueue) Push(val interface{}) *list.Element {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	e := q.list.PushFront(val)
	return e
}

func (q *GLockQueue) Pop() *list.Element {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	e := q.list.Back()
	q.list.Remove(e)
	return e
}
