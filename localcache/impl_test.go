package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_KEY = "test_key"
	TEST_VAL = "test_val"
)

func TestCacheGet(t *testing.T) {
	lastAccess := time.Now().Unix() - (CACHE_EXPIRED_SEC - 1)
	cache := cache{
		data: make(map[string]item),
	}
	cache.data[TEST_KEY] = item{
		val:     TEST_VAL,
		lastMod: lastAccess,
	}

	actual := cache.Get(TEST_KEY)
	assert.Equal(t, TEST_VAL, actual)
}

func TestExpiredCacheGet(t *testing.T) {
	lastAccess := time.Now().Unix() - (CACHE_EXPIRED_SEC + 1)
	cache := cache{
		data: make(map[string]item),
	}
	cache.data[TEST_KEY] = item{
		val:     TEST_VAL,
		lastMod: lastAccess,
	}

	actual := cache.Get(TEST_KEY)
	assert.Nil(t, actual)
}

func TestCacheSet(t *testing.T) {
	cache := cache{
		data: make(map[string]item),
	}
	cache.Set(TEST_KEY, TEST_VAL)

	actual := cache.data[TEST_KEY].val
	assert.Equal(t, TEST_VAL, actual)
}
