package _case

import (
	"fmt"
	"sync"
)

type onceMap struct {
	sync.Once
	data map[string]int
}

func (m *onceMap) LoadData() {
	//m.Once.Do()
	// once是协程单例的，仅加载一次
	m.Do(func() {
		list := []string{"A", "B", "C", "D"}
		for _, item := range list {
			_, ok := m.data[item]
			if !ok {
				m.data[item] = 0
			}
			m.data[item] += 1
		}
	})

	/*
		list := []string{"A", "B", "C", "D"}
		for _, item := range list {
			_, ok := m.data[item]
			if !ok {
				m.data[item] = 0
			}
			m.data[item] += 1
		}
	*/
}

func OnceCase() {
	o := &onceMap{
		data: make(map[string]int),
	}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.LoadData()
		}()
	}
	wg.Wait()
	fmt.Println(o.data)
}
