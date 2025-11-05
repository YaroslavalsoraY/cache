package rediscache

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	TestAddr = "localhost:6379"
	TestTTL = 10 * time.Second
	TestDataSize = 10
)

func TestNewRedisCache(t *testing.T) {
	cache := NewRedisCache(TestAddr, TestTTL)

	require.Equal(t, cache.ttl, TestTTL)
}

func TestSetGet(t *testing.T) {
	cache := NewRedisCache(TestAddr, TestTTL)

	for i := range TestDataSize {
		err := cache.Set(strconv.Itoa(i), strconv.Itoa(i))
		require.Equal(t, nil, err)
	}

	for i := range TestDataSize {
		result, err := cache.Get(strconv.Itoa(i))
		require.Equal(t, nil, err)
		require.Equal(t, strconv.Itoa(i), result)
	}
}

func TestExists(t *testing.T) {
	cache := NewRedisCache(TestAddr, TestTTL)

	for i := range TestDataSize {
		err := cache.Set(strconv.Itoa(i), strconv.Itoa(i))
		require.Equal(t, nil, err)
	}

	for i := range TestDataSize * 2 {
		result := cache.Exists(strconv.Itoa(i))
		if i < TestDataSize {
			require.Equal(t, result, true)
		} else {
			require.Equal(t, result, false)
		}
	}
}

func TestCount(t *testing.T) {
	cache := NewRedisCache(TestAddr, TestTTL)

	for i := range TestDataSize {
		cache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	count := cache.Count()

	require.Equal(t, count, TestDataSize)
}

func TestDelete(t *testing.T) {
	cache := NewRedisCache(TestAddr, TestTTL)

	for i := range TestDataSize {
		cache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range TestDataSize * 2 {
		isCounted := cache.Delete(strconv.Itoa(i))
		_, err := cache.Get(strconv.Itoa(i))
		require.NotEqual(t, err, nil)

		if i >= TestDataSize {
			require.Equal(t, isCounted, 0)
		}
	}

	require.Equal(t, cache.Count(), 0)
}