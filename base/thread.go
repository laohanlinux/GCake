package base

// thread obj, but i think need not to do, later i will improve, but now ......
import (
	"sync/atomic"

	//_ "GCake/net"

	"github.com/funny/goid"
	"github.com/laohanlinux/go-logger/logger"
)

type ThreadFunc func(...interface{}) interface{}
type Thread struct {
	func_       ThreadFunc
	c           chan string
	name_       string
	numCreated_ int32
	started_    bool
	signalbool  bool
	g_id        int32
}

func NewThread(threadFunc_ ThreadFunc, name string, join bool) *Thread {
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
func (t *Thread) Start() {
	if t.started_ == false {
		go t.startThread()
	} else {
		logger.Error("Failed in pthread_create")
		return
	}
}
func (t *Thread) Join() string {
	exit := <-t.c
	return exit
}
func (t *Thread) NumCreated() int32 {
	return t.numCreated_
}

func (t *Thread) SetSignalbool(b bool) {
	t.signalbool = b
}

func (t *Thread) startThread() {
	// update gid
	t.g_id = goid.Get()
	t.runInThread()
}

func (t *Thread) runInThread() {
	defer func() {
		if e := recover(); e != nil {
			logger.Error("sub thread runtime error: ", e)
		}
		if t.signalbool == true {
			t.c <- "normal"
			logger.Info("sent signalbool normal, sub_name is ", t.name_)
		}
		goruntineStore.c <- CurrentGoroutineId()
	}()
	logger.Info("start excute goroutine function body", t.name_, t.func_)
	t.func_(t.name_)
}
