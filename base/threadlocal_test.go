package base

import (
	"fmt"
	//"github.com/bmizerany/assert"
	"github.com/funny/goid"
	"runtime"
	"testing"
)

func t1(c chan int, i int) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
		c <- i
	}()
	var value GoruntineStoreData
	value = goid.Get() //100
	GoruntineSetSpecific(value)
}

func Test_ThreadLocal(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan int)
	for i := 0; i < 90; i++ {
		go t1(c, i)
	}
	for i := 0; i < 90; i++ {
		a := <-c
		fmt.Println(a)
	}
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	fmt.Printf("%s", buf)
	fmt.Println("main Pro exit")
	/* for k, v := range goruntineStore.gsVector {*/
	//assert.Equal(t, k, v)
	/*}*/
}
