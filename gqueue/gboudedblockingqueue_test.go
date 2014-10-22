package gqueue

import (
	"fmt"
	"runtime"
	"testing"
)

const ThreadNums = 100

func productGBoundedBlockingQueue(c chan<- int, i int, GCQ *GBoudedBlockingQueue) {
	defer func() { c <- 0 }()
	GCQ.Put(i)
	fmt.Println("Product: ", i)
}

func customGBoundedBlockingQueue(c chan<- int, GCQ *GBoudedBlockingQueue) {
	defer func() { c <- 0 }()
	fmt.Println("custom: ", GCQ.take().Value)
}

func Test_GBoundedBlockingQueue(t *testing.T) {
	c := make(chan int)
	GCQ := NewGBoundedBlockingQueue(ThreadNums)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := ThreadNums; i > 0; i-- {
		go customGBoundedBlockingQueue(c, GCQ)
	}

	for i := ThreadNums; i > 0; i-- {
		go productGBoundedBlockingQueue(c, i, GCQ)
	}

	for i := 2 * ThreadNums; i > 0; i-- {
		<-c
	}
	fmt.Println("queue size is: ", GCQ.size())
}
