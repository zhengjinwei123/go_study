package util

import (
	"container/list"
	"sync"
)

type keyValue struct {
	key   string
	value interface{}
}

type LRUCache struct {
	itemList *list.List
	itemMap  map[string]*list.Element
	maxSize  int
	lock     sync.Mutex
}

func (cache *LRUCache) Get(key string) interface{} {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	elem, ok := cache.itemMap[key]
	if !ok {
		return nil
	}
	cache.itemList.MoveToFront(elem)
	kv := elem.Value.(*keyValue)
	return kv.value
}

func (cache *LRUCache) Put(key string, val interface{}, timeout int64) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	elem, ok := cache.itemMap[key]

	if ok {
		// 找到了，更新记录
		cache.itemList.MoveToFront(elem)
		kv := elem.Value.(*keyValue)
		kv.value = val
	} else {
		// 是新元素，添加到列表
		elem := cache.itemList.PushFront(&keyValue{key: key, value: val})
		cache.itemMap[key] = elem

		// 超过上限，将最久未使用的移出
		if cache.itemList.Len() > cache.maxSize {
			delElem := cache.itemList.Back()
			kv := delElem.Value.(*keyValue)
			cache.itemList.Remove(delElem)
			delete(cache.itemMap, kv.key)
		}
	}
}

func (cache *LRUCache) Delete(key string) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	elem, ok := cache.itemMap[key]
	if ok {
		cache.itemList.Remove(elem)
		delete(cache.itemMap, key)
	}
}

func (cache *LRUCache) IsExist(key string) bool {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	if _, ok := cache.itemMap[key]; ok {
		return true
	}
	return false
}

func (cache *LRUCache) ClearAll() {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	for k, e := range cache.itemMap {
		cache.itemList.Remove(e)
		delete(cache.itemMap, k)
	}
}

func (cache *LRUCache) Len() int {
	return cache.itemList.Len()
}

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		itemList: list.New(),
		itemMap:  make(map[string]*list.Element),
		maxSize:  maxSize,
	}
}
