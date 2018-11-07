package singleton

import "sync"

type SingletonInitFunc func() (interface{},error)

type Singleton interface {
	Get() (interface{},error)
}

func NewSingleton(init SingletonInitFunc) Singleton {
	return &singletonImpl{init:init}
}

type singletonImpl struct {
	sync.Mutex
	data interface{}
	init SingletonInitFunc
	initialized bool
}

func (s *singletonImpl) Get() (interface{},error) {
	if s.initialized {
		return s.data,nil
	}

	s.Lock()
	defer s.Unlock()

	if s.initialized {
		return s.data,nil
	}

	var err error
	s.data,err = s.init()
	if err != nil {
		return nil,err
	}

	s.initialized = true
	return s.data,nil
}