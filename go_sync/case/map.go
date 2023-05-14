package _case

import (
	"fmt"
	"sync"
)

func MapCase() {
	/*
		正常情况下用法
			mp := make(map[string]int)
			mp["name"] = "nick"
	*/
	mp := sync.Map{}
	// 设置键值对
	mp.Store("name", "nick")
	mp.Store("email", "nick@voice.com")

	// 通过key获取value，如果不存在则返回nil,ok 否则返回false
	fmt.Println(mp.Load("name"))
	fmt.Println(mp.Load("email"))

	// 通过key获取value，如果不存在则设置指定的value
	// ok 为 true表示key存在并返回值，为false表示key 不存在并设置后返回
	fmt.Println(mp.LoadOrStore("hobby", "篮球"))
	fmt.Println(mp.LoadOrStore("hobby", "羽毛球"))

	// 根据key 获取value后，删除该key
	// ok 为 true表示key存在，为false表示key不存在
	fmt.Println(mp.LoadAndDelete("hobby"))
	fmt.Println(mp.LoadAndDelete("hobby"))

	// 为集合设置迭代函数，将为集合中的每一个键值对顺序调用该函数，如果该函数返回false，则停止迭代
	// 为遍历集合中所有元素提供便利
	mp.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

func MapCase1() {
	mp := sync.Map{}
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, item := range list {
				value, ok := mp.Load(item)
				if !ok {
					value, _ = mp.LoadOrStore(item, 0)
				}
				val := 0
				val = value.(int)
				val += 1
				mp.Store(item, val)
			}
		}()
	}
	wg.Wait()
	mp.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
	fmt.Println(mp)
}
