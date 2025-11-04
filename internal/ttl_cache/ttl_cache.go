package ttlcache

import (
	"errors"
	"sync"
	"time"
)

const DefaultTTL = 10 * time.Second

type TTLMap struct {
	cache map[string]TTLElement
	ttl   time.Duration
	mu    sync.RWMutex
}

type TTLElement struct {
	value    string
	removeAt time.Time
}

func NewTTLMap(ttl time.Duration) *TTLMap {
	if ttl < time.Second {
		ttl = DefaultTTL
	}
	newMap := TTLMap{cache: make(map[string]TTLElement), ttl: ttl, mu: sync.RWMutex{}}

	go newMap.mapCleaner()

	return &newMap
}

func (tm *TTLMap) Get(key string) (string, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	result, ok := tm.cache[key]
	if !ok {
		return "", errors.New("key does not exists")
	}

	return result.value, nil
}

func (tm *TTLMap) Set(key, value string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	newElement := TTLElement{
		value: value,
		removeAt: time.Now().Add(tm.ttl),
	}

	tm.cache[key] = newElement

	return nil
}

func (tm *TTLMap) Exists(key string) bool {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	_, ok := tm.cache[key]

	return ok
}

func (tm *TTLMap) Count() int {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	return len(tm.cache)
}

func (tm *TTLMap) Delete (keys ...string) int {
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

func (tm *TTLMap) mapCleaner() {
	for {
		<-time.Tick(tm.ttl)
		for k, v := range tm.cache {
			if v.removeAt.Before(time.Now()) {
				tm.mu.Lock()
				delete(tm.cache, k)
				tm.mu.Unlock()
			}
		}
	}
}