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

func add(args ...interface{}) interface{} {
	return 1 + 3
}

func Err(args ...interface{}) interface{} {
	fmt.Println(args[90])
	return nil
}
func Test_ThreadPool(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// create a thraed pool and it waits for sub thread exit
	T := NewThreadPool("MainThreadPool", true)
	// create 5 thraeds
	T.start(1)
	time.Sleep(1000000000)
	//put 2 thread
	f := print
	T.run(f)
	T.run(f)
	T.run(f)
	T.run(Err)
	t1 := add
	T.run(t1)
	T.run(f)

	time.Sleep(3000000000)
	for i := 0; i < 30000; i++ {
		if i/2 == 1 {
			T.run(f)
		} else {
			T.run(t1)
		}
	}
	T.stop()
}
