// localcache module for local cache implemented by map
package localcache

import "time"

const (
	// cacheExpiredSec define expiration time of cache
	cacheExpiredSec = 30
)

var (
	timeNow = time.Now
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
	now := timeNow().Unix()
	if data, ok := c.data[key]; ok && now-data.lastMod <= cacheExpiredSec {
		return data.val
	}
	return nil
}

// Set key value pair to cache
func (c *cache) Set(key string, val interface{}) {
	now := timeNow().Unix()
	c.data[key] = item{
		val:     val,
		lastMod: now,
	}
}
