package base

import (
	"github.com/laohanlinux/go-logger/logger"
	"runtime"
	"testing"
)

func threadFunc(args ...interface{}) interface{} {
	Instance()
	logger.Info(Instance(), " name=", Instance().Name())
	Instance().SetName("Only one, changed")
	return nil
}
func Test_SingleTon(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Instance().SetName("Only One")
	thread := NewThread(threadFunc, "hello", true)
	thread.Start()
	thread.Join()
	logger.Info(Instance(), Instance().Name())

}
