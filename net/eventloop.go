package net

import (
	"os"
)

type EventLoop struct {
	looping bool // atomic , 是否处于事件循环
	//断言是否在当前进程
	isInLoopThread bool
}

func (elp EventLoop) assertInLoopThread() {
	if !elp.isInLoopThread {
		os.Exit(127)
	}
}
