package net

import (
	. "GCake/base"
	. "GCake/tool"

	"github.com/laohanlinux/go-logger/logger"
)

type EventLoop struct {
	// atomic , 是否处于事件循环
	looping_ bool
	//当前对象所属的ID
	threadid_ int32

	t_looInThisThread *EventLoop
	this              *EventLoop
}

func NewEventLoop() *EventLoop {
	if GoruntineGetSpecific() != nil {
		logger.Fatal("Another EventLoop ", GoruntineGetSpecific(), " exists in this thread ", CurrentGoroutineId())
	}
	e := &EventLoop{
		looping_:  false,
		threadid_: CurrentGoroutineId(),
	}
	e.this = e
	GoruntineSetSpecific(e.threadid_)
	logger.Info("EventLoop created ", e, " In thread ", e.threadid_)
	return e
}

func (elp *EventLoop) Loop() {
	//事件循环，该函数不能跨线程电泳，只能在创建该对象的线程中调用
	//断言线程还没有事件循环
	Assert(!elp.looping_)
	//断言当前处于创建该对象的线程中
	elp.assertInLoopThread()
	elp.looping_ = true
	logger.Info("EventLoop ", elp, " start looping_")
	//
	// 事件循环标记设为false
	elp.looping_ = false
}

func (elp *EventLoop) IsInLoopThread() bool {
	return CurrentGoroutineId() == elp.threadid_
}

// 断言是否是当前线程
func (elp *EventLoop) assertInLoopThread() {
	if !elp.IsInLoopThread() {
		elp.abortNotInLoopThread()
	}
}

//终止线程
func (elp *EventLoop) abortNotInLoopThread() {
	logger.Fatal("EventLoop.assertInLoopThread EventLoop: ", &elp,
		"was created in thread_=", elp.threadid_, ", current thread_id=", elp.threadid_)
}

func (elp EventLoop) getEventLoopOfCurrentThread() *EventLoop {
	return elp.t_looInThisThread
}
