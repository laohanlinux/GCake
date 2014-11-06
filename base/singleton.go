package base

import (
	"sync"
)

var singleton *SingleTon

type SingleTon struct {
	name string
}

func Instance() *SingleTon {
	if singleton == nil {
		singleton = new(SingleTon)
	}
	return singleton
}

func init() {
	f := func() {
		Instance()
	}
	var once sync.Once
	once.Do(f)
}

func (s *SingleTon) SetName(name string) {
	s.name = name
}

func (s *SingleTon) Name() string {
	return s.name
}

func (s *SingleTon) destroy() {
	s = nil
}
