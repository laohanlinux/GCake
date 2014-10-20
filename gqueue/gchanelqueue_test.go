package gqueue

import (
	"fmt"
	//"github.com/bmizerany/assert"
	"reflect"
	"testing"
)

func Test_Chanels(t *testing.T) {
	GCQ := NewChanelQueue()
	fmt.Printf("%T\n", GCQ)
	GCQ.Push("a")
	GCQ.Push("b")
	GCQ.Push("c")
	fmt.Println(GCQ.Pop().Value)
	fmt.Println(reflect.ValueOf(GCQ.Pop().Value))
	fmt.Println(reflect.ValueOf(GCQ.Pop().Value))
	//assert.Equal(t, string(GCQ.Pop().Value), "a")
}
