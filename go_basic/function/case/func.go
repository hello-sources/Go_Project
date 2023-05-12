package _case

import (
	"errors"
	"log"
)

func Sum(a, b int) (sum int, err error) {
	if a <= 0 && b <= 0 {
		err = errors.New("两数相加要求两数不能同时小于等于0")
		return 0, err
	}
	sum = a + b
	return
}

// SumFunc 将函数作为类型
type SumFunc func(a, b int) (int, error)

// LogMiddleware 将函数作为输入输出实现中间件
func LogMiddleware(in SumFunc) SumFunc {
	// 返回的函数为闭包函数，其中 in 为闭包函数使用的外部函数内部变量
	return func(a, b int) (int, error) {
		log.Printf("日志中间件，记录操作数：a : %d, b : %d\n", a, b)
		return in(a, b)
	}
}

// Accumulation 声明 receiver 为函数类型的方法，即函数类型的对象的方法
func (sum SumFunc) Accumulation(list ...int) (int, error) {
	s := 0
	var err error
	for _, i := range list {
		s, err = sum(s, i)
		if err != nil {
			return s, err
		}
	}
	return s, err
}
