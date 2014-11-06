package base

import (
	"fmt"
	"github.com/laohanlinux/go-logger/logger"
	"testing"
)

func threadFunc_(args ...interface{}) interface{} {
	if len(args) == 0 {
		fmt.Println("run in thread function!!!!")
	} else {
		fmt.Println(args)
	}
	return nil
}

func Test_ThreadObj(t *testing.T) {
	logger.Info("Good!!!!")
	T := NewThread(threadFunc_, "one", true)
	T.Start()
	fmt.Println("waitting for thread exit")
	fmt.Println(T.Join())

	T = NewThread(threadFunc_, "one", false)
	T.Start()

	T = NewThread(threadFunc_, "one", true)
	T.Start()
	fmt.Println("waitting for thread exit")
	fmt.Println(T.Join())
}
