package cache

import (
	"sync"
	"time"

	"github.com/mailgun/groupcache"
)

type GroupCache struct {
	group *groupcache.Group
	mu    sync.Mutex
}

func NewGroupCache(name string, cacheBytes int64) *GroupCache {
	return &GroupCache{
		group: groupcache.NewGroup(name, cacheBytes, groupcache.GetterFunc(
			func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
				// Return nil as the value, indicating that the value does not exist in the cache
				return nil
			},
		)),
	}
}

func (gc *GroupCache) Set(key string, value interface{}, d time.Duration) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	// Use the key as the cache item's string value
	//item := groupcache.StringCacheItem(value.(string))
	//
	//// Set the value in the groupcache group
	//gc.group.Set(nil, key, item)

	// Schedule the item for removal after the specified duration
	time.AfterFunc(d, func() {
		gc.Delete(key)
	})
}

func (gc *GroupCache) Get(key string) (interface{}, bool) {
	var data string
	// Get the value from the groupcache group
	err := gc.group.Get(nil, key, groupcache.StringSink(&data))
	if err != nil {
		return nil, false
	}
	return data, true
}

func (gc *GroupCache) Delete(key string) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	// Delete the value from the groupcache group
	//gc.group.Delete(nil, key)
}

func (gc *GroupCache) Clear() {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	// Clear all values from the groupcache group
	//gc.group.Clear()
}

func (gc *GroupCache) Size() int {
	// Return the number of items in the groupcache group
	//return gc.group.CacheStats().Gets
	return 0
}
