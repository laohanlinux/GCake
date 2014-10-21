package gqueue

import (
	"sync"
)

type MutexLock struct {
	mutex sync.Mutex
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

func (MLG MutexLockGuard) NewMutexLock(mutex *MutexLock) {
	(*mutex).lock()
}

func (MLG MutexLockGuard) DeleteMutexLock(mutex *MutexLock) {
	(*mutex).unlock()
}
