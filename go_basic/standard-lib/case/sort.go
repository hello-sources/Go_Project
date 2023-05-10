package _case

import (
	"fmt"
	"sort"
)

type sortUser struct {
	ID   int
	Name string
	Age  uint8
}

type ByID []sortUser

// Len 获取长度
func (a ByID) Len() int {
	return len(a)
}

// Swap 交换位置
func (a ByID) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less 返回i位置的ID是否小于j位置的ID
func (a ByID) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}

func SortCase() {
	list := []sortUser{
		{ID: 10, Name: "nick", Age: 18},
		{ID: 8, Name: "aick", Age: 19},
		{ID: 11, Name: "ncck", Age: 17},
		{ID: 20, Name: "dick", Age: 11},
		{ID: 1, Name: "fick", Age: 28},
		{ID: 3, Name: "fidk", Age: 128},
		{ID: 18, Name: "dnck", Age: 8},
		{ID: 21, Name: "aiek", Age: 12},
		{ID: 6, Name: "nidk", Age: 17},
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Age < list[j].Age
	})
	fmt.Println(list)

	// 实现sort.Interface接口
	sort.Sort(ByID(list))
	fmt.Println(list)
}
