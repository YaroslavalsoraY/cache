package cache

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Exists(key string) bool
	Count() int
	Delete (keys ...string) int
}