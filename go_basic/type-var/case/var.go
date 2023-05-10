package _case

import "fmt"

func VarDeclareCase() {
	// 通过var声明变量
	var i, z, x int = 1, 2, 3
	var j int = 100

	// 通过var关键词声明变量，并赋值
	var f float32 = 100.123

	// 通过 := 以推断的方式定义变量并赋值，此方式只能使用局部变量定义
	b := true

	// 数组,长度不可变
	var arr = [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{2, 3, 4, 5}
	var arr2 [5]int
	arr2[2] = 4
	arr2[3] = 5
	fmt.Println(i, z, x, j, f, b, arr, arr1, arr2)

	// 指针类型,用于表示变量地址的类型
	var intPtr *int
	var floatPtr *float64
	var i1 = 100
	f1(&i1)
	fmt.Println(intPtr, floatPtr, i1)

	// 接口类型
	var inter interface{}
	fmt.Println(inter)
	inter = i1
	fmt.Println(inter)
}

func f1(i *int) {
	*i = *i + 1
}
