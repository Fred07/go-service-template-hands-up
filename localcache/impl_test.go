package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testKey = "test_key"
	testVal = "test_val"
)

func TestCacheGet(t *testing.T) {
	lastAccess := time.Now().Unix() - (cacheExpiredSec - 1)
	cache := cache{
		data: make(map[string]item),
	}
	cache.data[testKey] = item{
		val:     testVal,
		lastMod: lastAccess,
	}

	actual := cache.Get(testKey)
	assert.Equal(t, testVal, actual)
}

func TestExpiredCacheGet(t *testing.T) {
	lastAccess := time.Now().Unix() - (cacheExpiredSec + 1)
	cache := cache{
		data: make(map[string]item),
	}
	cache.data[testKey] = item{
		val:     testVal,
		lastMod: lastAccess,
	}

	actual := cache.Get(testKey)
	assert.Nil(t, actual)
}

func TestCacheSet(t *testing.T) {
	timeNow = func() time.Time { return time.Unix(1629446406, 0) }

	cache := cache{
		data: make(map[string]item),
	}
	cache.Set(testKey, testVal)

	actual := cache.data[testKey].val
	actualLastMod := cache.data[testKey].lastMod
	assert.Equal(t, testVal, actual)
	assert.Equal(t, timeNow().Unix(), actualLastMod)
}
