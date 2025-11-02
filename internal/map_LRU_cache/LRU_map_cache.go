package LRUMapCache

import (
	linkedList "cache/internal/map_LRU_cache/linked_list"
	"errors"
	"sync"
)

const DEFAULTCAPACITY = 8

type LRUMapCache struct {
	ll *linkedList.LinkedList
	cache map[string]*linkedList.Element
	capacity int
	mu *sync.RWMutex
}

func NewLRUMapCache(capacity int) *LRUMapCache {
	if capacity < 1 {
		capacity = DEFAULTCAPACITY
	}
	
	return &LRUMapCache{
		ll: linkedList.NewLinkedList(), 
		cache: make(map[string]*linkedList.Element, capacity),
		capacity: capacity,
		mu: &sync.RWMutex{},	
	}
}

func (mc *LRUMapCache) Get(key string) (string, error) {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	
	result, ok := mc.cache[key]
	if !ok {
		return "", errors.New("key does not exists")
	}

	mc.ll.MoveToFirst(result)
	return result.GetValue(), nil
}

func (mc *LRUMapCache) Set(key, value string) error {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	
	result, ok := mc.cache[key]
	if ok {
		mc.ll.DeleteElement(result)
	}

	if mc.ll.GetLen() == mc.capacity {
		lastKey := mc.ll.GetLastKey()
		delete(mc.cache, lastKey)
		mc.ll.DeleteLast()
	}

	newData := mc.ll.AddFirst(key, value)
	mc.cache[key] = newData
	
	return nil
}

func (mc *LRUMapCache) Exists(key string) bool {
	mc.mu.RLock()
	defer mc.mu.RUnlock()

	_, ok := mc.cache[key]

	return ok
}

func (mc *LRUMapCache) Count() int {
	mc.mu.RLock()
	defer mc.mu.RUnlock()

	return len(mc.cache)
}

func (mc *LRUMapCache) Delete(keys ...string) int {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	
	var count int
	for _, key := range keys {
		result, ok := mc.cache[key]
		
		if ok {
			mc.ll.DeleteElement(result)
			delete(mc.cache, key)
			count++
		}
	}

	return count
}