package ttlcache

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	TestTTL      = 3 * time.Second
	TestDataSize = 10
)

func TestNewTTLMap(t *testing.T) {
	newMap := NewCacheWithTTL(TestTTL)

	newDefaultMap := NewCacheWithTTL(time.Microsecond)

	require.Equal(t, newDefaultMap.ttl, DefaultTTL)

	require.Equal(t, newMap.ttl, TestTTL)
}

func TestSetGet(t *testing.T) {
	ttlMap := NewCacheWithTTL(TestTTL)

	for i := 0; i < TestDataSize; i++ {
		ttlMap.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := 0; i < TestDataSize; i++ {
		el, err := ttlMap.Get(strconv.Itoa(i))
		require.Equal(t, err, nil)
		require.Equal(t, el, strconv.Itoa(i))
	}

	el, err := ttlMap.Get(strconv.Itoa(TestDataSize + 1))
	require.NotEqual(t, err, nil)
	require.Equal(t, el, "")
}

func TestExists(t *testing.T) {
	ttlMap := NewCacheWithTTL(TestTTL)

	for i := range TestDataSize {
		ttlMap.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TestDataSize * 2 {
		result := ttlMap.Exists(strconv.Itoa(i))
		if i < TestDataSize {
			require.Equal(t, result, true)
		} else {
			require.Equal(t, result, false)
		}
	}
}

func TestCount(t *testing.T) {
	ttlMap := NewCacheWithTTL(TestTTL)

	for i := range TestDataSize {
		ttlMap.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	count := ttlMap.Count()

	require.Equal(t, count, TestDataSize)
}

func TestDelete(t *testing.T) {
	ttlMap := NewCacheWithTTL(TestTTL)

	for i := range TestDataSize {
		ttlMap.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TestDataSize * 2 {
		isCounted := ttlMap.Delete(strconv.Itoa(i))
		_, err := ttlMap.Get(strconv.Itoa(i))
		require.NotEqual(t, err, nil)

		if i >= TestDataSize {
			require.Equal(t, isCounted, 0)
		}
	}

	require.Equal(t, ttlMap.Count(), 0)
}

func TestMapCleanerAllMap(t *testing.T) {
	ttlMap := NewCacheWithTTL(TestTTL)

	for i := range TestDataSize {
		ttlMap.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	<-time.Tick(TestTTL + 1*time.Second)

	require.Equal(t, 0, ttlMap.Count())
}

func TestMapCleanerNotAllMap(t *testing.T) {
	ttlMap := NewCacheWithTTL(TestTTL)

	for i := range TestDataSize {
		ttlMap.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	<-time.Tick(TestTTL / 2)

	ttlMap.Set(strconv.Itoa(TestDataSize), strconv.Itoa(TestDataSize))

	<-time.Tick(TestTTL/2 + 1*time.Second)

	result := ttlMap.Exists(strconv.Itoa(TestDataSize))

	require.Equal(t, result, true)
	for i := range TestDataSize {
		result = ttlMap.Exists(strconv.Itoa(i))
		require.Equal(t, false, result)
	}
}
