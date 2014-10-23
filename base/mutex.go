package base

import (
	"sync"
)

type MutexLock struct {
	mutex sync.Mutex
}

func NewMutexLock() *MutexLock {
	m_ := new(MutexLock)
	return m_
}
func (m MutexLock) isLockByThisThread() {}

func (m MutexLock) assertLocked() {}

func (m MutexLock) lock() {
	m.mutex.Lock()
}

func (m MutexLock) unlock() {
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
	(*mutex_).lock()
	defer (*mutex_).unlock()
	return f()
}

func (MLG MutexLockGuard) NewMutexLock(mutex *MutexLock) {
	(*mutex).lock()
}

func (MLG MutexLockGuard) DeleteMutexLock(mutex *MutexLock) {
	(*mutex).unlock()
}
