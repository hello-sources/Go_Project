package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发协程不需要通讯，可以使用 sync.WaitGroup来实现
var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	wg.Done()
}

var ch = make(chan string, 10)

func testChannel(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url
}

func main() {
	for i := 0; i < 3; i++ {
		// wg.add添加一个计数，wg.Done减去一个计数
		wg.Add(1)
		// go + func启动一个协程，并发执行func，且协程间不需要通讯
		go download("a.com/" + string(i+'0'))
	}
	wg.Wait()
	fmt.Println("Done!")

	// 通过channel接收返回信息
	for i := 0; i < 3; i++ {
		go testChannel("b.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println("finish", msg)
	}
	fmt.Println("Finsh channel 传输!")

	return
}
