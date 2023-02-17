package main

import (
	"errors"
	"fmt"
)

type student struct {
	name string
	age  int
}

type person interface {
	getname() string
	getGender() string
}

type worker struct {
	name string
	age  int
}

type cleaner struct {
	name   string
	gender string
}

func (work *worker) getname() string {
	return work.name
}

func (clean *cleaner) getname() string {
	return clean.name
}

func (clean *cleaner) getGender() string {
	return clean.gender
}

func add(num1, num2 int) int {
	return num1 + num2
}

func div(num1, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error, name is null")
	}
	fmt.Println("hello, ", name)
	return nil
}

func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error发生，剩下的交给defer和recover")
			ret = -1
		}
	}()
	arr := [3]int{1, 2, 3}
	return arr[index]
}

// 方法前加上实例名，可以通过实例名访问该实例的字段name和其他方法了
func (stu *student) helloStruct(person string) string {
	return fmt.Sprintf("hello %s, I am %s\n", person, stu.name)
}

func main() {
	quo, rem := div(100, 7)
	fmt.Println(quo, rem)
	fmt.Println(quo + rem)

	// error.New返回自定义错误
	if err := hello(""); err != nil {
		fmt.Println(err)
	}

	// defer和recover
	// defer用来处理异常，只要发生异常，在协程退出前，会执行defer挂载的任务
	// 如果触发panic，就会执行defer
	fmt.Println(get(5), "\nfinished")

	// 结构体与method
	stu := &student{
		name: "Rose",
	}
	msg := stu.helloStruct("Jack")
	fmt.Print(msg)

	// 接口，接口定义了一组方法的集合，接口不能被实例化，一个类型可以实现多个接口。
	// 必须实现接口里面的所有方法
	var c person = &cleaner{
		name:   "保洁阿姨",
		gender: "female",
	}
	fmt.Println(c.getname(), c.getGender())

	// var p person = &worker{
	// 	name: "lalala",
	// 	age:  19,
	// }
	// st := p.(*worker)
	// fmt.Println(st.getname())

	// 空接口
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 99, 85}
	fmt.Println(m)
}
