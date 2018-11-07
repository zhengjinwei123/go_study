package util

import (
	"errors"
	"sync"
	"newsCaptors/util/singleton"
)

var Factory fatory
func init(){
	Factory.instances = make(map[string]singleton.Singleton)
}

type fatory struct {
	instances map[string]singleton.Singleton
	lock sync.Mutex
}

func (f *fatory) Set(name string , init singleton.SingletonInitFunc) bool {
	f.lock.Lock()
	defer f.lock.Unlock()
	if _,ok := f.instances[name]; !ok {
		f.instances[name] = singleton.NewSingleton(init)
		return true
	}
	return false
}

func (f *fatory) Get(name string) (interface{},error) {
	if _,ok := f.instances[name];ok {
		return f.instances[name].Get()
	}
	return nil,errors.New("fatory get error:"+name+" not exists")
}

