package gqueue

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_gqueue(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	gbq := NewGBlockingQueue()
	for i := 10; i > 0; i-- {
		go func() {
			fmt.Println("wait queue!!!!")
			worker := gbq.take()
			fmt.Printf("worker: %T, %v\n", worker, worker)
		}()
	}

	for i := 0; i < 10; i++ {
		gbq.Put(i)
		fmt.Println("put: ", i)
		time.Sleep(time.Duration(100 * 1000 * 1000))
	}
}
