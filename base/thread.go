package base

// thread obj, but i think need not to do, later i will improve, but now ......
import (
	"fmt"
	"sync/atomic"
	//"github.com/bmizerany/assert"
)

type Thread struct {
	func_       func(...interface{}) interface{}
	c           chan string
	name_       string
	numCreated_ int32
	started_    bool
	signalbool  bool
}

func NewThread(threadFunc_ func(...interface{}) interface{}, name string, join bool) *Thread {
	started_ := false
	func_ := threadFunc_
	name_ := name

	var numCreated_ int32 = 0
	if join == false {
		return &Thread{
			func_:       func_,
			name_:       name_,
			numCreated_: atomic.AddInt32(&numCreated_, 1),
			started_:    started_,
			signalbool:  join,
		}
	} else {
		c := make(chan string)
		return &Thread{
			func_:       func_,
			name_:       name_,
			c:           c,
			numCreated_: atomic.AddInt32(&numCreated_, 1),
			started_:    started_,
			signalbool:  join,
		}
	}
}
func (t *Thread) start() {
	if t.started_ == false {
		go t.startThread()
	} else {
		fmt.Println("Failed in pthread_create")
		return
	}
}
func (t *Thread) join() string {
	exit := <-t.c
	return exit
}
func (t *Thread) numCreated() int32 {
	return t.numCreated_
}

func (t *Thread) SetSignalbool(b bool) {
	t.signalbool = b
}

func (t *Thread) startThread() {
	t.runInThread()
}

func (t *Thread) runInThread() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("sub thread runtime error: ", e)
		}
		if t.signalbool == true {
			t.c <- "normal"
			fmt.Println("sent signalbool normal, sub_name is ", t.name_)
		}
	}()
	t.func_(t.name_)
}
