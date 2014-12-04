package net

import (
	. "GCake/base"
	"fmt"
	"testing"
)

/*测试EventLoop是否是一个Thread一个EventLoop*/

func threadFunc(arg ...interface{}) interface{} {
	fmt.Println("threadFunc(): gid = ", CurrentGoroutineId())
	loop := NewEventLoop()
	loop.Loop()
	return nil
}

func Test_EventLoop(t *testing.T) {
	fmt.Println("main gid = ", CurrentGoroutineId())
	loop := NewEventLoop()

	t1 := NewThread(threadFunc, "t1", true)
	t1.Start()
	loop.Loop()
	t1.Join()

	loop1 := NewEventLoop()
	loop1.Loop()
}
