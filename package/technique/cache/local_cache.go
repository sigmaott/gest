package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type GoCache struct {
	cache *cache.Cache
}

func NewGoCache(defaultExpiration, cleanupInterval time.Duration) *GoCache {
	return &GoCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (gc *GoCache) Set(key string, value any, d time.Duration) {
	gc.cache.Set(key, value, d)
}

func (gc *GoCache) Get(key string) (interface{}, bool) {
	value, found := gc.cache.Get(key)
	return value, found
}

func (gc *GoCache) Delete(key string) {
	gc.cache.Delete(key)
}

func (gc *GoCache) Clear() {
	gc.cache.Flush()
}

func (gc *GoCache) Size() int {
	return gc.cache.ItemCount()
}
