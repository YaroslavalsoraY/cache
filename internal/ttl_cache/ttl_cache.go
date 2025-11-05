package ttlcache

import (
	"errors"
	"sync"
	"time"
)

const DefaultTTL = 10 * time.Second

type CacheWithTTL struct {
	cache map[string]ElementWithTTL
	ttl   time.Duration
	mu    sync.RWMutex
}

type ElementWithTTL struct {
	value    string
	removeAt time.Time
}

func NewCacheWithTTL(ttl time.Duration) *CacheWithTTL {
	if ttl < time.Second {
		ttl = DefaultTTL
	}
	newMap := CacheWithTTL{cache: make(map[string]ElementWithTTL), ttl: ttl, mu: sync.RWMutex{}}

	go newMap.mapCleaner()

	return &newMap
}

func (tm *CacheWithTTL) Get(key string) (string, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	result, ok := tm.cache[key]
	if !ok || result.removeAt.Before(time.Now()) {
		return "", errors.New("key does not exists")
	}

	return result.value, nil
}

func (tm *CacheWithTTL) Set(key, value string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	newElement := ElementWithTTL{
		value:    value,
		removeAt: time.Now().Add(tm.ttl),
	}

	tm.cache[key] = newElement

	return nil
}

func (tm *CacheWithTTL) Exists(key string) bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	_, ok := tm.cache[key]

	return ok
}

func (tm *CacheWithTTL) Count() int {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	return len(tm.cache)
}

func (tm *CacheWithTTL) Delete(keys ...string) int {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	var count int

	for _, key := range keys {
		if _, ok := tm.cache[key]; ok {
			delete(tm.cache, key)
			count++
		}
	}

	return count
}

func (tm *CacheWithTTL) mapCleaner() {
	for {
		<-time.Tick(tm.ttl)
		tm.mu.Lock()
		for k, v := range tm.cache {
			if v.removeAt.Before(time.Now()) {
				delete(tm.cache, k)
			}
		}
		tm.mu.Unlock()
	}
}
