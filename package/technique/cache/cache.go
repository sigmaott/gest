package cache

import "time"

type ICache interface {
	// Set a value in the cache with an associated key
	Set(key string, value any, d time.Duration)

	// Get a value from the cache with a given key
	Get(key string) (any, bool)

	// Delete a value from the cache with a given key
	Delete(key string)

	// Clear all values from the cache
	Clear()

	// Get the number of items in the cache
	Size() int
}
