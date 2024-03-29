package _case

import (
	"fmt"
	"sync"
	"time"
)

func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})
	go ReadList(&list, cond)
	go ReadList(&list, cond)
	go ReadList(&list, cond)
	time.Sleep(time.Second * 10)
	initList(&list, cond)

}

func initList(list *[]int, c *sync.Cond) {
	// 主叫方，可以持有锁，也可以不持有锁
	c.L.Lock()
	defer c.L.Unlock()
	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}

	//c.Broadcast() // 唤醒所有条件等待的协程
	c.Signal() // 只唤醒其中一个协程
}

func ReadList(list *[]int, c *sync.Cond) {
	// 被叫方，必须持锁
	c.L.Lock()
	defer c.L.Unlock()
	for len(*list) == 0 {
		fmt.Println("readlist wait")
		c.Wait()
	}
	fmt.Println("list 数据为：", *list)
}

type queue struct {
	list []int
	cond *sync.Cond
}

func newQueue() *queue {
	q := &queue{
		list: []int{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
	return q
}

func (q *queue) Put(item int) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.list = append(q.list, item)
	// 当数据写入成功后，唤醒一个协程来处理数据
	q.cond.Signal()
}

func (q *queue) GetMany(n int) []int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.list) < n {
		q.cond.Wait()
	}
	list := q.list[:n]
	q.list = q.list[n:]
	return list
}

func CondQueueCase() {
	q := newQueue()
	var wq sync.WaitGroup
	for n := 1; n <= 10; n++ {
		wq.Add(1)
		go func(n int) {
			defer wq.Done()
			list := q.GetMany(n)
			fmt.Printf("%d : %d \n", n, list)
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
	}
	wq.Wait()
}
