## Thread And Thread Pool 

这一章是线程对象和线程池的设计，虽然在`go`中，线程是非常低资源的，但是为了和`muduo`一样，暂时先这样，以后再改吧。


- Thread

**结构设计：**

```
type Thread struct {
        func_       func(...interface{})
        c           chan string
        name_       string
        numCreated_ int32
        started_    bool
        signalbool  bool
}
```

`func_`为线程函数体，这是基于对象的设计方式，其实就是一个用户自定的回调函数；`c`是一个外部线程和该线程的单向通信的通道，主要用来模拟`posix`的`join`; `name_`线程的名字；`numCreated32`备用；`started`用于表明线程是否已启动；`signalbool`是否返回退出信号。

线程执行流程：


`start() -> startThread() -> runInThread() -> func_() -> send siganl`,其实有效的执行体是`runThread()`里的`func_()`。

**源码实现：**
```
func NewThread(threadFunc_ func(...interface{}), name string, join bool) Thread {
        started_ := false
        func_ := threadFunc_
        name_ := name

        var numCreated_ int32 = 0
        if join == false {
                return Thread{
                        func_:       func_,
                        name_:       name_,
                        numCreated_: atomic.AddInt32(&numCreated_, 1),
                        started_:    started_,
                        signalbool:  join,
                }
        } else {
                c := make(chan string)
                return Thread{
                        func_:       func_,
                        name_:       name_,
                        c:           c,
                        numCreated_: atomic.AddInt32(&numCreated_, 1),
                        started_:    started_,
                        signalbool:  join,
                }
        }
}
func (t Thread) start() {
        if t.started_ == false {
                go t.startThread()
        } else {
                fmt.Println("Failed in pthread_create")
                return
        }
}
func (t Thread) join() string {
        exit := <-t.c
        return exit
}
func (t Thread) numCreated() int32 {
        return t.numCreated_
}

func (t Thread) startThread() {
        t.runInThread()
}

func (t Thread) runInThread() {
        defer func() {
                if e := recover(); e != nil {
                        fmt.Println(e)
                }
        }()
        t.func_()
        if t.signalbool == true {
                t.c <- "normal"
        }
}
```

- threadpool

