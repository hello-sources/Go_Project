package main

import (
	"context"
	"fmt"
	_case "goBasic/function/case"
	"os"
	"os/signal"
)

func main() {
	// 调用函数
	fmt.Println(_case.Sum(-1, 10))
	// 将函数赋值给变量
	f1 := _case.Sum
	// 执行函数
	fmt.Println(f1(1, 2))

	// 将函数作为输入输出作为中间件
	f2 := _case.LogMiddleware(_case.Sum)
	// 再次附加中间件
	f2 = _case.LogMiddleware(f2)
	fmt.Println(f2(1, 2))

	f3 := _case.SumFunc(f1)
	fmt.Println(f3.Accumulation(1, 2, 3, 4))
	fmt.Println(f2.Accumulation(1, 2, 3, 4, 5))

	fmt.Println(_case.Fib(10))

	// 闭包的陷阱
	_case.ClosureTrap()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	<-ctx.Done()
}
