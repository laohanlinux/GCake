package base

import (
	"container/list"
	"fmt"
	"strconv"
)

type ThreadPool struct {
	// task, task should be a function, the function is about handle some thing of caculating, store or others,
	// it also should be a function for thread
	mutex_   *MutexLock
	cond_    *Condition
	name_    string
	threads_ []*Thread
	running_ bool
	queue_   *list.List // queue should store task, and the task  should be a function
	join_    bool
}

func NewThreadPool(name string, join bool) *ThreadPool {
	mutex_ := NewMutexLock()
	cond_ := NewCondition(mutex_.mutex)
	name_ := name
	running_ := false
	queue_ := list.New()
	return &ThreadPool{
		mutex_:   mutex_,
		cond_:    cond_,
		name_:    name_,
		running_: running_,
		join_:    join,
		queue_:   queue_,
	}
}

// start thread pool, that is say that start thread function
func (t *ThreadPool) start(numThreads int) {
	if len(t.threads_) != 0 {
		fmt.Println("ThreadPool start fail because of threads is not empty!!!")
		return
	}
	t.running_ = true
	for i := 0; i < numThreads; i++ {
		fmt.Println("start ", i, " thraed ")
		t.threads_ = append(t.threads_, NewThread(t.runInThread, strconv.Itoa(i), t.join_))
		// start thread
		t.threads_[i].start()
	}
	fmt.Println("finish all subthread start ....", len(t.threads_))
}

func (t *ThreadPool) stop() {
	f := func(...interface{}) interface{} {
		t.running_ = false
		t.cond_.notifyAll()
		return nil
	}
	LockAndUnlock(t.mutex_, f)
	if t.join_ {
		for _, k := range t.threads_ {
			fmt.Println("wait sub exit")
			k.join()
		}
	}
}

func (t *ThreadPool) run(task func(args ...interface{}) interface{}) {
	if len(t.threads_) == 0 {
		fmt.Println("handle in main threads, sub num is ", len(t.threads_))
		task()
	} else {
		LockAndUnlock(t.mutex_, func(...interface{}) interface{} {
			fmt.Printf("push a task in queue and the task is %T\n", task)
			t.queue_.PushFront(task)
			t.cond_.notify()
			return nil
		})
	}
}

func (t *ThreadPool) take() interface{} {
	t.mutex_.lock()
	defer func() {
		if t.mutex_ != nil {
			t.mutex_.unlock()
		}
	}()
	//if pool is stop, it also jump out the block
	for t.queue_.Len() == 0 && t.running_ {
		fmt.Println("queue is empty, condition waiting task")
		t.cond_.wait()
	}
	if t.queue_.Len() > 0 {
		e := t.queue_.Back()
		t.queue_.Remove(e)
		return e.Value
	} else {
		fmt.Println("Queue is empty!!NONONO")
		return nil
	}
}

func (t *ThreadPool) runInThread(args ...interface{}) interface{} {
	defer func() {
		panic("sub abort error")
		/* if e := recover(); e != nil {*/
		//// if we want to disaster recovery for sub thread, can do that
		//t.mutex_.lock()
		//[>if len(args) > 0 {<]
		////switch n := args[0].(type) {
		////case string:
		////if i, err := strconv.Atoi(n); err == nil {
		////t.threads_[i].SetSignalbool(false)
		////t.threads_[i] = NewThread(t.runInThread, n, t.join_)
		////t.threads_[i].start()
		////}
		////default:
		////panic("args must be sub thread name")
		////}
		//[>}<]
		//panic(e)
		//t.mutex_.unlock()
		/*}*/
	}()
	// thread pool should be running
	for t.running_ {
		task := t.take()
		if task == nil {
			if t.running_ == false {
				fmt.Println("the pool is exit, so sub thread also exit\n")
			} else {
				fmt.Println("the queue should not be empty!!!")
			}
		} else {
			switch tk := task.(type) {
			case func(...interface{}) interface{}:
				tk()
			default:
				fmt.Println("task format is error")
			}
		}
	}
	return nil
}
