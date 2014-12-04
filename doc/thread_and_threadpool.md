## Thread And Thread Pool

这一章是线程对象和线程池的设计，虽然在`go`中，线程是非常低资源的，但是为了和`muduo`一样，暂时先这样，以后再改吧。


- ### Thread ###

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

- ### threadpool ###

这个线程池主要功能:

   - 主线程（线程池线程）在线程池为空的时候，会自己执行任务。

   - 子线程只有在线程池处于运行状态时才可以执行任务

   - 线程池停止时，会通知子线程，子线程接收到信息会退出，线程池线程可以获取子线程的退出信息

   - 当线程池发生意外退出时，不会影响到其他线程；但是线程池线程不会创建新的线程来替代意外退出的线程

   - 线程池是通过信号量来实现的

   - 任务是一个多参数函数体

   - 任务队列是一个线程安全的`list`容器

   -     `LockAndUnlock` 函数可以简单实现锁，因为`golang`没有构造函数和虚函数，所以使用`defer`简单实现


线程池的主要成员有：
```
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
```

**TODO**

- 简化任务函数

- 循环队列替代`list`

- 超时任务队列


