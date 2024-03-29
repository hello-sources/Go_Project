package _case

import (
	"context"
	"fmt"
	"time"
)

func ContextCase() {
	ctx := context.Background()                         // 创建context对象
	ctx = context.WithValue(ctx, "desc", "ContextCase") // 给context对象创建值
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	data := [][]int{{1, 2}, {3, 2}}
	ch := make(chan []int)
	go calculate(ctx, ch)
	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}
	time.Sleep(10 * time.Second)
}

func calculate(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			ctx := context.WithValue(ctx, "desc", "calculate")

			ch := make(chan []int)
			go sumContext(ctx, ch)
			ch <- item

			ch1 := make(chan []int)
			go multiContext(ctx, ch1)
			ch1 <- item

			//fmt.Println(item)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("calculate 协程退出，context desc : %s, 错误信息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func sumContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := sum(a, b)
			fmt.Printf("%d + %d = %d \n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("sumContext 协程退出，context desc : %s, 错误信息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func multiContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := multi(a, b)
			fmt.Printf("%d * %d = %d \n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("multiContext 协程退出，context desc : %s, 错误信息：%s", desc, ctx.Err())
			return
		}
	}
}

func sum(a, b int) int {
	return a + b
}

func multi(a, b int) int {
	time.Sleep(5 * time.Second)
	return a * b
}
