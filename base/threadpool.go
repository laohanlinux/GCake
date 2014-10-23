package base

import (
	"container/list"
	"fmt"
	"strconv"
)

type ThreadPool struct {
	// task
	task_    func(...interface{}) interface{}
	mutex_   *MutexLock
	cond_    *Condition
	name_    string
	threads_ []Thread
	running_ *bool
	queue_   *list.List // queue should store task, and the task  should be a function
	join_    bool
}

func NewThreadPool(name string, join bool) *ThreadPool {
	mutex_ := NewMutexLock()
	cond_ := NewCondition(&mutex_.mutex)
	name_ := name
	running_ := new(bool)
	*running_ = false
	queue_ := list.New()
	return &ThreadPool{
		mutex_:   mutex_,
		cond_:    cond_,
		name_:    name_,
		running_: running_,
		join_:    join,
		task_:    threadpoolFunc, //define sub thread fyunction
		queue_:   queue_,
	}
}

func threadpoolFunc(i ...interface{}) interface{} {
	if len(i) > 0 {
		switch t := i[0].(type) {
		case ThreadPool:
			LockAndUnlock(t.mutex_, func(args ...interface{}) interface{} {
				for t.queue_.Len() == 0 && *t.running_ {
					t.cond_.wait()
				}
				if t.queue_.Len() > 0 {
					e := t.queue_.Back()
					t.queue_.Remove(e)
					return e.Value
				} else {
					return nil
				}
			})
		default:

		}
	} else {
		fmt.Println("threadpoolFunc excute fail")
	}
	return nil
}

// start thread pool, that is say that start thread function
func (t ThreadPool) start(numThreads int) {
	if len(t.threads_) != 0 {
		fmt.Println("ThreadPool start fail because of threads is not empty!!!")
		return
	}
	*t.running_ = true
	for i := 0; i > numThreads; i++ {
		t.threads_ = append(t.threads_, NewThread(t.task_, strconv.Itoa(i), t.join_))
		// start thread
		t.threads_[i].start()
	}
}

func (t ThreadPool) stop() {
	f := func(...interface{}) interface{} {
		*t.running_ = false
		t.cond_.notifyAll()
		// wether waiting for sub thread exit signal
		if t.join_ {
			for _, k := range t.threads_ {
				k.join()
			}
		}
		return nil
	}
	LockAndUnlock(t.mutex_, f)
}

func (t ThreadPool) runInThread() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	// thread pool should be running
	for *t.running_ {
		task := t.task_()
		switch tk := task.(type) {
		case func(...interface{}) interface{}:
			tk()
		default:
		}
	}
}
