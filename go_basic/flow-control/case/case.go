package _case

import (
	"fmt"
)

func FlowControlCase() {
	ifElseCase(1)
	ifElseCase(2)
	ifElseCase(3)
	ifElseCase(0)

	forCase()

	switchCase("A", 1)
	switchCase("C", "ch")
	switchCase("W", 1.2)

	gotoCase()
}

func ifElseCase(a int) {
	if a == 0 {
		fmt.Println("执行if语句")
	} else if a == 1 {
		fmt.Println("执行else if 语句")
	} else {
		fmt.Println("执行else语句")
	}
}

func forCase() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("位置1执行for语句块 i : ", i)
	}

	var i int
	var b = true
	for b {
		fmt.Println("位置2执行for语句块 i : ", i)
		i++
		if i >= 10 {
			b = false
		}
	}

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for ind, data := range list {
		fmt.Printf("位置3执行for语句块 index = %d, data = %d\n", ind, data)
	}

	mp := map[string]string{"A": "a", "B": "b", "C": "c"}
	for key, value := range mp {
		fmt.Printf("位置4执行for语句块 key = %s, value = %s\n", key, value)
	}

	str := "今天天气很好"
	for ind, ch := range str {
		fmt.Printf("位置5执行for语句块 index：%v, char : %v\n", ind, string(ch))
	}

	var j int
	for {
		fmt.Println("位置6执行for语句块j:", j)
		j++
		if j >= 10 {
			break
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("位置7执行for语句块, i = %d, j = %d\n", i, j)
		}
	}

lab1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("位置8执行for语句块, 跳出循环, i = %d, j = %d\n", i, j)
			break lab1
		}
	}
}

func switchCase(alpha string, in interface{}) {
	switch alpha {
	case "A":
		fmt.Println("执行A语句块")
	case "B":
		fmt.Println("执行B语句块")
	case "C", "D":
		fmt.Println("执行C、D语句块")
		fallthrough
	case "E":
		fmt.Println("执行E语句块")
	case "F":
		fmt.Println("执行F语句块")
	}

	switch in.(type) {
	case string:
		fmt.Println("in 输入为字符串")
	case int:
		fmt.Println("in 输入为int类型")
	default:
		fmt.Println("未能识别in的类型")
	}
}

func gotoCase() {
	var a = 0
lab1:
	fmt.Println("goto位置1")
	for i := 0; i < 10; i++ {
		a += i
		if a == 0 {
			a += 1
			goto lab1
		}
	}
}
