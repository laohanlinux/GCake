package base

import (
	"fmt"
	"github.com/jtolds/gls"
	"testing"
)

func Test_ThreadLocal(t *testing.T) {
	fmt.Println("ok")
	var testObj1 *ThreadLocal = NewThreadLocal()
	var testObj2 *ThreadLocal = NewThreadLocal()
	var threadId = gls.GenSym()
	tob := gls.Values{threadId: "12345"}
	testObj1.SetValue(tob)
	tob[threadId] = "909090"
	testObj2.SetValue(tob)

}
