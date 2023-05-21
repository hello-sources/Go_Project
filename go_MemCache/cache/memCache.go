package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type memCache struct {
	//最大内存
	maxMemorySize int64
	//最大内存字符串表示
	maxMemorySizeStr string
	//当前已使用内存
	currentMemorySize int64
	// 缓存键值对
	values map[string]*memCacheValue
	// 读写锁，可以并发读，只能互斥写
	locker sync.RWMutex
	// 清除过期缓存时间间隔
	clearExpiredItemTimeInterval time.Duration
}

type memCacheValue struct {
	// value值
	val interface{}
	// 过期时间
	expireTime time.Time
	// 有效时长
	expire time.Duration
	// value 大小
	size int64
}

func NewMemCache() Cache {
	mc := &memCache{
		values:                       make(map[string]*memCacheValue, 0),
		clearExpiredItemTimeInterval: time.Second,
	}
	go mc.clearExpiredItem()
	return mc
}

func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	return true
}

func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValueSize(val),
	}
	mc.del(key)
	mc.add(key, v)
	if mc.currentMemorySize > mc.maxMemorySize {
		mc.del(key)
		log.Fatal(fmt.Sprintf("max memory size : %d", mc.maxMemorySize))
		panic(fmt.Sprintf("max memory size : %d", mc.maxMemorySize))
	}
	return true
}

func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}

func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		mc.currentMemorySize -= tmp.size
		delete(mc.values, key)
	}
}

func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currentMemorySize += val.size
}

func (mc *memCache) Get(key string) (interface{}, bool) {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	mcv, ok := mc.get(key)
	if ok {
		// 判断缓存是否过期
		if mcv.expire != 0 && mcv.expireTime.Before(time.Now()) {
			mc.del(key)
			return nil, false
		}
		return mcv.val, ok
	}
	return nil, false
}

func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}

func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	_, ok := mc.values[key]
	return ok
}

func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.values = make(map[string]*memCacheValue, 0)
	mc.currentMemorySize = 0
	return true
}

func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	return int64(len(mc.values))
}

func (mc *memCache) clearExpiredItem() {
	timeTicker := time.NewTicker(mc.clearExpiredItemTimeInterval)
	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			for key, item := range mc.values {
				if item.expire != 0 && time.Now().After(item.expireTime) {
					mc.locker.Lock()
					mc.del(key)
					mc.locker.Unlock()
				}
			}
		}
	}
}
