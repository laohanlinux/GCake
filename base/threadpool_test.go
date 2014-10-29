package base

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func print(args ...interface{}) interface{} {
	fmt.Println("Hello Word", time.Now().Unix(), " task: ", args)
	return nil
}
func Test_ThreadPool(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// create a thraed pool and it waits for sub thread exit
	T := *NewThreadPool("MainThreadPool", true)
	// create 5 thraeds
	T.start(5)
	fmt.Println("wait ...")
	time.Sleep(1000000000)
	//put 2 thread
	f := print
	T.run(f)
	T.run(f)

	time.Sleep(3000000000)
	for i := 0; i < 5; i++ {
		T.run(f)
	}
	T.stop()
}
