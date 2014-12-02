package base

import (
	"github.com/jtolds/gls"
)

type TObj interface{}

type LocalObj gls.Values

type TSetFunc func()

type ThreadLocal struct {
	mgr *gls.ContextManager
}

func NewThreadLocal() *ThreadLocal {
	return &ThreadLocal{
		mgr: gls.NewContextManager(),
	}
}

func (tl *ThreadLocal) Value(threadId gls.ContextKey) (TObj, bool) {
	return tl.mgr.GetValue(threadId)
}

func (tl *ThreadLocal) SetValue(new_values gls.Values) { //context_call TSetFunc) {
	/*if context_call == nil {*/
	//context_call = func() {}
	/*}*/
	tl.mgr.SetValues(new_values, func() {})
}

func (tl *ThreadLocal) GetValue(threadId gls.ContextKey) (TObj, bool) {
	return tl.Value(threadId)
}
