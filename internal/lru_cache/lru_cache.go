package lrucache

import (
	linkedList "cache/internal/lru_cache/linked_list"
	"errors"
	"sync"
)

const DefaultCapacity = 8

type LRUCache struct {
	ll       *linkedList.LinkedList
	cache    map[string]*linkedList.Element
	capacity int
	mu       sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity < 1 {
		capacity = DefaultCapacity
	}

	return &LRUCache{
		ll:       linkedList.NewLinkedList(),
		cache:    make(map[string]*linkedList.Element, capacity),
		capacity: capacity,
		mu:       sync.RWMutex{},
	}
}

func (mc *LRUCache) Get(key string) (string, error) {
	mc.mu.RLock()
	defer mc.mu.RUnlock()

	result, ok := mc.cache[key]
	if !ok {
		return "", errors.New("key does not exists")
	}

	mc.ll.MoveToFirst(result)
	return result.GetValue(), nil
}

func (mc *LRUCache) Set(key, value string) error {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	result, ok := mc.cache[key]
	if ok {
		mc.ll.Delete(result)
	}

	if mc.ll.GetLen() == mc.capacity {
		lastKey := mc.ll.DeleteLast().GetKey()
		delete(mc.cache, lastKey)
	}

	newData := mc.ll.AddFirst(key, value)
	mc.cache[key] = newData

	return nil
}

func (mc *LRUCache) Exists(key string) bool {
	mc.mu.RLock()
	defer mc.mu.RUnlock()

	_, ok := mc.cache[key]

	return ok
}

func (mc *LRUCache) Count() int {
	mc.mu.RLock()
	defer mc.mu.RUnlock()

	return len(mc.cache)
}

func (mc *LRUCache) Delete(keys ...string) int {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	var count int
	for _, key := range keys {
		result, ok := mc.cache[key]

		if ok {
			mc.ll.Delete(result)
			delete(mc.cache, key)
			count++
		}
	}

	return count
}
