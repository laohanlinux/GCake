package glockqueue

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
	return &ChanelQueue{mutex, list}
}

func (q *GLockQueue) Size() int {
	mutex.Lock()
	defer mutex.Unlock()
	return q.list.Len()
}

func (q *GLockQueue) push(val interface{}) *list.Element {
	mutex.Lock()
	defer mutex.Unlock()
	e := q.list.PushFront(val)
	return e
}

func (q *GLockQueue) Pop() *list.Element {
	mutex.Lock()
	defer mutex.Unlock()
	e := q.list.Back()
	q.list.Remove(e)
	return e
}
