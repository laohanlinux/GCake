package base

import (
	"fmt"
	"testing"
)

func threadFunc(args ...interface{}) {
	if len(args) == 0 {
		fmt.Println("run in thread function!!!!")
	} else {
		fmt.Println(args)
	}
}

func threadFunc1(a, b int) int {
	return a + b
}

func Test_ThreadObj(t *testing.T) {
	T := NewThread(threadFunc, "one", true)
	T.start()
	fmt.Println("waitting for thread exit")
	fmt.Println(T.join())

	T = NewThread(threadFunc, "one", false)
	T.start()

	T = NewThread(threadFunc, "one", true)
	T.start()
	fmt.Println("waitting for thread exit")
	fmt.Println(T.join())
}
