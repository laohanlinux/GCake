package base

import (
	"fmt"
	"sync"
)

type MutexLock struct {
	mutex *sync.Mutex
}

func NewMutexLock() *MutexLock {
	this := new(MutexLock)
	this.mutex = new(sync.Mutex)
	return this
}

func (m MutexLock) isLockByThisThread() {}

func (m MutexLock) assertLocked() {}

func (m *MutexLock) lock() {
	m.mutex.Lock()
}

func (m *MutexLock) unlock() {
	m.mutex.Unlock()
}

func (m MutexLock) getPThreadMutex() *MutexLock {
	return &m
}

////////////
type MutexLockGuard struct {
	mutex *MutexLock
}

/// too complex .....
func LockAndUnlock(mutex_ *MutexLock, f func(args ...interface{}) interface{}) interface{} {
	mutex_.lock()
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	f()
	mutex_.unlock()
	return nil
}

func (MLG MutexLockGuard) NewMutexLock(mutex *MutexLock) {
	(*mutex).lock()
}

func (MLG MutexLockGuard) DeleteMutexLock(mutex *MutexLock) {
	(*mutex).unlock()
}
