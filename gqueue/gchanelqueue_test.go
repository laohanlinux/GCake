package gqueue

import (
	"fmt"
	"github.com/bmizerany/assert"
	"reflect"
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
	c := GCQ.Pop().Value
	fmt.Printf("c: %T\n", c)
	fmt.Println(reflect.ValueOf(GCQ.Pop().Value))
	fmt.Println(reflect.ValueOf(GCQ.Pop().Value))
	q := Q{"huarong", 23}
	GCQ.Push(q)
	q1 := q.Value(GCQ.Pop().Value)
	var a interface{} = Q{"name", 89}
	fmt.Printf("a: %T\n", a)
	assert.Equal(t, q.name, q1.name)
}
