package base

import (
	"sync"
)

type Condition struct {
	cond *sync.Cond
}

func NewCondition(lock ...*sync.Mutex) *Condition {
	var mutext *sync.Mutex
	if len(lock) == 0 {
		mutext = new(sync.Mutex)
	} else {
		mutext = lock[0]
	}
	newcond := sync.NewCond(mutext)
	return &Condition{newcond}
}

func (c *Condition) wait() {
	c.cond.Wait()
}

func (c *Condition) notify() {
	c.cond.Signal()
}

func (c *Condition) notifyAll() {
	c.cond.Broadcast()
}
