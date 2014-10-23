package base

import (
	"container/list"
	"fmt"
	"strconv"
)

type ThreadPool struct {
	// task
	Task     func(...interface{})
	mutex_   *MutexLock
	cond_    Condition
	name_    string
	threads_ []Thread
	running_ bool
	queue_   list.List
	join_    bool
}

func NewThreadPool(name string, join bool) *ThreadPool {
	mutex_ := NewMutexLock()
	cond_ := NewCondition(&mutex_.mutex)
	name_ := name
	running_ := false
	return &ThreadPool{
		mutex_:   mutex_,
		cond_:    *cond_,
		name_:    name_,
		running_: running_,
		join_:    join,
	}
}

func (t ThreadPool) start(numThreads int) {
	if len(t.threads_) != 0 {
		fmt.Println("ThreadPool start fail because of threads is not empty!!!")
		return
	}
	t.running_ = true
	for i := 0; i > numThreads; i++ {
		t.threads_ = append(t.threads_, NewThread(t.Task, strconv.Itoa(i), t.join_))
		// start thread
		t.threads_[i].start()
	}
}

func (t ThreadPool) stop() {
	f := func(...interface{}) {
		t.running_ = false
		t.cond_.notifyAll()
		// wether waiting for sub thread exit signal
		if t.join_ {
			for _, k := range t.threads_ {
				k.join()
			}
		}
	}
	LockAndUnlock(t.mutex_, f)
}
