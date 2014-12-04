package base

import (
	"github.com/funny/goid"
)

/*type CurrentThread struct {*/
////线程真实的pid(tid)的缓存, muduo lib 为了防止多次调用
////syscall(SYS_getpid)来获取pid，所以做了一个缓存量，
////但是golang的线程是一种携程，所以就不用了
//t_cachedTid  int
//t_threadName string
/*}*/

func CurrentGoroutineId() int32 {
	return goid.Get()
}
