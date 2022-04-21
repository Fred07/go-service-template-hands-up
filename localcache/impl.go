// localcache module for local cache implemented by map
package localcache

import "time"

const (
	// CACHE_EXPIRED_SEC define expiration time of cache
	CACHE_EXPIRED_SEC = 30
)

type item struct {
	val     interface{}
	lastMod int64
}

type cache struct {
	data map[string]item
}

// Get cache data by key string
func (c *cache) Get(key string) interface{} {
	now := time.Now().Unix()
	if data, ok := c.data[key]; ok && now-data.lastMod <= CACHE_EXPIRED_SEC {
		return data.val
	}
	return nil
}

// Set key value pair to cache
func (c *cache) Set(key string, val interface{}) {
	now := time.Now().Unix()
	c.data[key] = item{
		val:     val,
		lastMod: now,
	}
}
