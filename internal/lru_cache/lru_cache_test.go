package lrucache

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

const TestCapacity = 1000

func TestNewLRUCache(t *testing.T) {
	mapCache := NewLRUCache(TestCapacity)
	mapCacheDefault := NewLRUCache(-TestCapacity)

	require.Equal(t, mapCache.capacity, TestCapacity)
	require.Equal(t, mapCacheDefault.capacity, DefaultCapacity)
}

func TestGetSet(t *testing.T) {
	mapCache := NewLRUCache(TestCapacity)

	for i := range TestCapacity {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TestCapacity {
		result, err := mapCache.Get(strconv.Itoa(i))
		require.Equal(t, err, nil)
		require.Equal(t, result, strconv.Itoa(i))
	}
}

func TestExists(t *testing.T) {
	mapCache := NewLRUCache(TestCapacity)

	for i := range TestCapacity {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TestCapacity * 2 {
		result := mapCache.Exists(strconv.Itoa(i))
		if i < TestCapacity {
			require.Equal(t, result, true)
		} else {
			require.Equal(t, result, false)
		}
	}
}

func TestCount(t *testing.T) {
	mapCache := NewLRUCache(TestCapacity)

	for i := range TestCapacity {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	result := mapCache.Count()

	require.Equal(t, result, TestCapacity)
}

func TestDelete(t *testing.T) {
	mapCache := NewLRUCache(TestCapacity)

	for i := range TestCapacity {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TestCapacity * 2 {
		isCounted := mapCache.Delete(strconv.Itoa(i))
		_, err := mapCache.Get(strconv.Itoa(i))
		require.NotEqual(t, err, nil)

		if i >= TestCapacity {
			require.Equal(t, isCounted, 0)
		}
	}

	require.Equal(t, mapCache.Count(), 0)
}

func TestLRUAlgorithm(t *testing.T) {
	mapCache := NewLRUCache(TestCapacity)

	for i := range mapCache.capacity {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	mapCache.Get(strconv.Itoa(0))
	mapCache.Set(strconv.Itoa(mapCache.capacity), strconv.Itoa(mapCache.capacity))

	exists := mapCache.Exists(strconv.Itoa(0))
	require.Equal(t, exists, true)

	exists = mapCache.Exists(strconv.Itoa(1))
	require.Equal(t, exists, false)
}
