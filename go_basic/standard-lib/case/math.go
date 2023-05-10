package _case

import (
	"fmt"
	"math"
)

func MathCase() {
	fmt.Println("2的10次方：", math.Pow(2, 10))
	fmt.Println("返回以2为底，1024的对数：", math.Log(1024))
	fmt.Println("返回两个数中较大的数：", math.Max(2, 10))
	fmt.Println("向上取整：", math.Ceil(2.12))
	fmt.Println("向下取整：", math.Floor(2.89))
	fmt.Println("90度角的正弦值：", math.Sin(math.Pi/2))
	fmt.Println("1 反正弦值：", math.Asin(1))

}
