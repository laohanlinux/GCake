package base

import (
	"fmt"
	"testing"
)

func Test_Mutex(t *testing.T) {
	m := NewMutexLock()
	f := func(...interface{}) interface{} {
		fmt.Println("hello Word")
		return nil
	}
	LockAndUnlock(m, f)

}
