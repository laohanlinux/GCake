package gqueue

import (
	"fmt"
	"github.com/bmizerany/assert"
	"reflect"
	"testing"
)

type QB struct {
	name string
	age  int
}

func (q QB) Value(i interface{}) QB {
	switch value := i.(type) {
	case QB:
		q = value
	default:
		fmt.Println(" value is error")
	}
	return q
}

func Test_GLockQueue(t *testing.T) {
	GCQ := NewGLockQueue()
	GCQ.Push("a")
	GCQ.Push("b")
	GCQ.Push("c")
	c := GCQ.Pop().Value
	fmt.Printf("c: %T\n", c)
	fmt.Println(reflect.ValueOf(GCQ.Pop().Value))
	fmt.Println(reflect.ValueOf(GCQ.Pop().Value))
	q := QB{"huarong", 23}
	GCQ.Push(q)
	q1 := q.Value(GCQ.Pop().Value)
	var a interface{} = QB{"name", 89}
	fmt.Printf("a: %T\n", a)
	assert.Equal(t, q.name, q1.name)
}
