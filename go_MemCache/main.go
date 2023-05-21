package main

import (
	"fmt"
	cache_server "memcache/cache-server"
	"time"
)

func main() {
	cache := cache_server.NewMemCache()
	cache.SetMaxMemory("300GB")
	cache.Set("int", 1, time.Second)
	cache.Set("bool", false, time.Second)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)
	cache.Set("int", 1)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1})
	cache.Get("int")
	cache.Del("int")
	cache.Flush()
	cache.Keys()
	fmt.Println(cache.Keys())

	cache.SetMaxMemory("300GB")
	cache.Set("int", 1, time.Second)
	cache.Set("bool", false, time.Second)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)
	cache.Set("int", 1)
	cache.Set("bool", false)
	cache.Set("data", map[string]interface{}{"a": 1})
	fmt.Println(cache.Get("int"))
	fmt.Println(cache.Get("bool"))
	fmt.Println(cache.Keys())

	//fmt.Println(cache.GetValueSize(1))
	//fmt.Println(cache.GetValueSize(false))
	//fmt.Println(cache.GetValueSize("测试"))
	//fmt.Println(cache.GetValueSize(map[string]string{
	//	"name": "测试",
	//	"addr": "测试测试再测试",
	//}))
	//
	//fmt.Println(unsafe.Sizeof(1))
	//fmt.Println(unsafe.Sizeof(false))
	//fmt.Println(unsafe.Sizeof("测试"))
	//fmt.Println(unsafe.Sizeof(map[string]string{
	//	"name": "测试",
	//	"addr": "测试测试再测试",
	//}))
}
