package _case

import "fmt"

func NewCase() {
	// 通过new函数，可以创建任意类型，并返回指针
	mpPtr := new(map[string]*user)
	if *mpPtr == nil {
		fmt.Println("map默认值为空，但是地址不为空")
	}
	//(*mpPtr)["A"] = &user{}
	fmt.Println(mpPtr)

	slicePtr := new([]user)
	if *slicePtr == nil {
		fmt.Println("切片值为nil", *slicePtr)
	}
	//fmt.Println((*slicePtr)[0])
	*slicePtr = append(*slicePtr, user{Name: "nick"})
	fmt.Println(slicePtr)

	userPtr := new(user)
	strPtr := new(string)
	fmt.Println(userPtr, strPtr)
}

// MakeCase ：make 仅用于切片、集合、通道的初始化
func MakeCase() {
	// 初始化切片，并设置长度和容量
	slice := make([]int, 10, 20)
	slice[0] = 10
	fmt.Println(slice)

	// 初始化集合，并设置集合的初始大小
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	fmt.Println(mp)

	// 初始化通道，设置通道的读写方向和缓冲大小
	ch := make(chan int, 10)
	chW := make(chan<- int, 10) // 只写
	chR := make(<-chan int, 10) // 只读，缓冲区可以设置也可以不设置
	fmt.Println(ch, chW, chR)
}

func SliceAndMapCase() {
	// 定义切片
	var slice []int
	slice = []int{1, 2, 3, 4, 5, 6}
	slice1 := make([]int, 10)
	slice1[1] = 10
	fmt.Println(slice, slice1)

	slice2 := make([]int, 5, 10)
	slice2[0] = 0
	slice2[1] = 1
	slice2[2] = 2
	slice2[3] = 3
	slice2[4] = 4
	fmt.Println(slice2, len(slice2), cap(slice2))

	// 切片的截取
	slice3 := slice2[1:3]
	fmt.Println(slice3, len(slice3), cap(slice3))
	slice4 := slice2[1:10]
	fmt.Println(slice4, len(slice4), cap(slice4))

	// 切片的附加
	slice4 = append(slice4, 11, 22, 33, 44)
	fmt.Println(slice4)

	// 集合，无序
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	mp["B"] = "b"
	mp["C"] = "c"
	mp["D"] = "d"
	mp["E"] = "e"
	fmt.Println(mp)
	// 无序
	for k, v := range mp {
		fmt.Println(k, v)
	}
	// 删除集合元素
	delete(mp, "B")
	fmt.Println(mp)
}
