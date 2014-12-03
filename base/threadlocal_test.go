package base

import (
	"fmt"
	"github.com/bmizerany/assert"
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

const n = 1000

func Test_ThreadLocal(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan int)
	for i := 0; i < n; i++ {
		go t1(c, i)
	}
	for i := 0; i < n; i++ {
		<-c
	}
	for key, value := range goruntineStore.gsVector {
		switch v := value.(type) {
		case int:
			assert.Equal(t, key, v)
		default:
		}
	}
}
