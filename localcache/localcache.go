// localcache module for local cache implemented by map
package localcache

// Cache is struct to provide interface of cache manipulation
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{})
}

// New to get Cache interface
func New() Cache {
	c := &cache{
		data: make(map[string]item),
	}
	return c
}
