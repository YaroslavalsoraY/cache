package rediscache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	ttl		time.Duration
}

func NewRedisCache(addr string, ttl time.Duration) *RedisCache {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
	 	panic(err)
	}

	return &RedisCache{client: rdb, ttl: ttl}
}

func (rc *RedisCache) Set(key, value string) error {
	ctx := context.Background()
	
	err := rc.client.Set(ctx, key, value, rc.ttl).Err()

	return err
}

func (rc *RedisCache) Get(key string) (string, error) {
	ctx := context.Background()

	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (rc *RedisCache) Exists(key string) bool {	
	ctx := context.Background()

	res, _ := rc.client.Exists(ctx, key).Result()

	if res == 1 {
		return true
	}

	return false
}

func (rc *RedisCache) Count() int {
	ctx := context.Background()

	res, _ := rc.client.DBSize(ctx).Result()

	return int(res)
}

func (rc *RedisCache) Delete(keys ...string) int {
	ctx := context.Background()

	keysDeleted, _ := rc.client.Del(ctx, keys...).Result()

	return int(keysDeleted)
}