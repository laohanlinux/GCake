package base

import (
	"fmt"
	"github.com/bmizerany/assert"
	//"reflect"
	"runtime"
	"testing"
)

type Q struct {
	name string
	age  int
}

func (q Q) Value(i interface{}) Q {
	switch value := i.(type) {
	case Q:
		q = value
	default:
		fmt.Println(" value is error")
	}
	return q
}

func Test_Chanels(t *testing.T) {
	GCQ := NewChanelQueue()
	GCQ.Push("a")
	GCQ.Push("b")
	GCQ.Push("c")
	GCQ.Pop()
	GCQ.Pop()
	GCQ.Pop()
	q := Q{"huarong", 23}
	GCQ.Push(q)
	q1 := q.Value(GCQ.Pop().Value)
	var a interface{} = Q{"name", 89}
	fmt.Printf("a: %T\n", a)
	assert.Equal(t, q.name, q1.name)
}

func product(c chan<- int, i int, GCQ *GChanelQueue) {
	defer func() { c <- 0 }()
	GCQ.Push(i)
	fmt.Println("Product: ", i)
}

func custom(c chan<- int, GCQ *GChanelQueue) {
	defer func() { c <- 0 }()
	fmt.Println("custom: ", GCQ.Pop().Value)
}

func Test_thread(t *testing.T) {
	c := make(chan int)
	GCQ := NewChanelQueue()
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 10; i > 0; i-- {
		go product(c, i, GCQ)
	}

	for i := 10; i > 0; i-- {
		go custom(c, GCQ)
	}

	for i := 20; i > 0; i-- {
		fmt.Println("receive: ", <-c)
	}
}
