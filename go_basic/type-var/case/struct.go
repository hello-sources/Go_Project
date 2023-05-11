package _case

import "fmt"

type user struct {
	Name string
	Age  uint
	Addr Address
}

type Address struct {
	Province string
	City     string
}

func StructCase() {
	// 值类型
	u := user{
		Name: "彭于晏",
		Age:  18,
	}
	f2(&u)
	fmt.Println(u)

	// 指针类型
	u1 := &user{
		Name: "nick",
		Age:  6,
	}
	fmt.Println(u1)
	f2(u1)
	fmt.Println(u1)

	// 指针类型
	u2 := new(user)
	fmt.Println("默认值：", u2)
	u2.Name = "nick3"
	u2.Age = 8
	fmt.Println(u2, u2.Addr.Province)

	// 结构体为值类型，定义变量后将会被默认初始化
	var u3 user
	fmt.Println(u3)

}

func f2(u *user) {
	u.Age = 28
}
