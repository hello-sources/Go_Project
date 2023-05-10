package _case

import "fmt"

func Simplecase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Println(getMaxNumInt(a, b))
	fmt.Println(getMaxNumFloat(c, d))

	fmt.Println(getMaxNum(a, b))
	fmt.Println(getMaxNum(c, d))
}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 加入interface包起来目的是为了避免出现*int指针类型被识别成乘法
func getMaxNum[T interface{ int | float64 }](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CusNumT interface {
	// ~表示支持类型的衍生类型
	// 多行之间取交集，单行内 | 取并集
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint16
}

// MyInt64 是int64的衍生类型，具有基础int64的新类型，与int64是不同类型
type MyInt64 int64

// MyInt32 是int32的别名，与int32是同一类型
type MyInt32 = int32

func CusNumTCase() {
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = a, b
	fmt.Println("自定义泛型：", getMaxCusNum(a, b))
	fmt.Println("自定义泛型：", getMaxCusNum(a1, b1))

	var c, d float64 = 5, 6
	fmt.Println("自定义泛型：", getMaxCusNum(c, d))

	var e, f int64 = 7, 8
	var g, h MyInt64 = 7, 8
	fmt.Println("自定义泛型：", getMaxCusNum(e, f))
	fmt.Println("自定义泛型：", getMaxCusNum(g, h))
}

func getMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 内置类型
func BuiltInCase() {
	var a, b string = "abc", "efg"
	fmt.Println("内置类型 comparable 泛型类型约束：", getBuiltInComparable(a, b))

	var c, d float64 = 500, 250
	fmt.Println("内置类型 comparable 泛型类型约束：", getBuiltInComparable(c, d))

	var f = 123.456
	printBuiltInAny(a)
	printBuiltInAny(f)
}

func getBuiltInComparable[T comparable](a, b T) bool {
	//comparable类型，只支持 == != 两个操作
	if a == b {
		return true
	}
	return false
}

func printBuiltInAny[T any](a T) {
	fmt.Println("内置类型 any泛型类型约束：", a)
}
