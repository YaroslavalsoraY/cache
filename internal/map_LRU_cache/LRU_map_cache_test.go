package LRUMapCache

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)
const TESTCAPACITY = 1000

func TestNewLRUMapCache(t *testing.T) {
	mapCache := NewLRUMapCache(TESTCAPACITY)
	mapCacheDefault := NewLRUMapCache(-TESTCAPACITY)

	require.Equal(t, mapCache.capacity, TESTCAPACITY)
	require.Equal(t, mapCacheDefault.capacity, DEFAULTCAPACITY)
}

func TestGetSet(t *testing.T) {
	mapCache := NewLRUMapCache(TESTCAPACITY)

	for i := range TESTCAPACITY {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TESTCAPACITY {
		result, err := mapCache.Get(strconv.Itoa(i))
		require.Equal(t, err, nil)
		require.Equal(t, result, strconv.Itoa(i))
	}
}

func TestExists(t *testing.T) {
	mapCache := NewLRUMapCache(TESTCAPACITY)

	for i := range TESTCAPACITY {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TESTCAPACITY * 2 {
		result := mapCache.Exists(strconv.Itoa(i))
		if i < TESTCAPACITY {
			require.Equal(t, result, true)
		} else {
			require.Equal(t, result, false)
		}
	}
}

func TestCount(t *testing.T) {
	mapCache := NewLRUMapCache(TESTCAPACITY)

	for i := range TESTCAPACITY {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	result := mapCache.Count()

	require.Equal(t, result, TESTCAPACITY)
}

func TestDelete(t *testing.T) {
	mapCache := NewLRUMapCache(TESTCAPACITY)

	for i := range TESTCAPACITY {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TESTCAPACITY * 2 {
		isCounted := mapCache.Delete(strconv.Itoa(i))
		_, err := mapCache.Get(strconv.Itoa(i))
		require.NotEqual(t, err, nil)

		if i >= TESTCAPACITY {
			require.Equal(t, isCounted, 0)
		}
	}

	require.Equal(t, mapCache.Count(), 0)
}

func TestLRU(t *testing.T) {
	mapCache := NewLRUMapCache(TESTCAPACITY)

	for i := range mapCache.capacity {
		mapCache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	mapCache.Get(strconv.Itoa(0))
	mapCache.Set(strconv.Itoa(mapCache.capacity), strconv.Itoa(mapCache.capacity))

	isExists := mapCache.Exists(strconv.Itoa(0))
	require.Equal(t, isExists, true)

	isExists = mapCache.Exists(strconv.Itoa(1))
	require.Equal(t, isExists, false)
}