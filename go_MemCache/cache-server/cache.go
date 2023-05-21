package cache_server

import (
	"memcache/cache"
	"time"
)

type cacheServer struct {
	memCache cache.Cache
}

func NewMemCache() *cacheServer {
	cache := &cacheServer{
		memCache: cache.NewMemCache(),
	}
	return cache
}

func (cs *cacheServer) SetMaxMemory(size string) bool {
	cs.memCache.SetMaxMemory(size)
	return true
}

func (cs *cacheServer) Set(key string, val interface{}, expire ...time.Duration) bool {
	expireTs := time.Second * 0
	if len(expire) > 0 {
		expireTs = expire[0]
	}
	return cs.memCache.Set(key, val, expireTs)
}

func (cs *cacheServer) Get(key string) (interface{}, bool) {
	return cs.memCache.Get(key)
}

func (cs *cacheServer) Del(key string) bool {
	return cs.memCache.Del(key)
}

func (cs *cacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}

func (cs *cacheServer) Flush() bool {
	return cs.memCache.Flush()
}

func (cs *cacheServer) Keys() int64 {
	return cs.memCache.Keys()
}
