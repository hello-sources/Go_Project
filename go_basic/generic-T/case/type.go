package _case

import "fmt"

type user struct {
	ID   int64
	Name string
	Age  uint8
}

type address struct {
	ID       int
	Province string
	City     string
}

// 集合转列表
func mapToList[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}

func myPrint[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypeCase() {
	userMp := make(map[int64]user, 0)
	userMp[1] = user{ID: 1, Name: "nick", Age: 18}
	userMp[2] = user{ID: 2, Name: "tom", Age: 19}
	userList := mapToList[int64, user](userMp)

	ch := make(chan user)
	go myPrint(ch)
	for _, u := range userList {
		ch <- u
	}

	addrMp := make(map[int]address, 0)
	addrMp[1] = address{ID: 1, Province: "安徽", City: "合肥"}
	addrMp[2] = address{ID: 2, Province: "江苏", City: "南京"}
	addrList := mapToList[int, address](addrMp)

	addrch := make(chan address)
	go myPrint(addrch)
	for _, a := range addrList {
		addrch <- a
	}
}

// List 定义一个泛型切片
type List[T any] []T

// MapT 泛型集合的定义
// 声明两个泛型，分别为k,v
type MapT[k comparable, v any] map[k]v

// Chan 泛型通道的定义
type Chan[T any] chan T

func TTypeCase1() {
	userMp := make(MapT[int64, user], 0)
	userMp[1] = user{ID: 1, Name: "nick", Age: 18}
	userMp[2] = user{ID: 2, Name: "tom", Age: 19}

	var userList List[user]
	userList = mapToList[int64, user](userMp)

	ch := make(Chan[user])
	go myPrint(ch)
	for _, u := range userList {
		ch <- u
	}

	addrMp := make(MapT[int, address], 0)
	addrMp[1] = address{ID: 1, Province: "安徽", City: "合肥"}
	addrMp[2] = address{ID: 2, Province: "江苏", City: "南京"}

	var addrList List[address]
	addrList = mapToList[int, address](addrMp)

	addrch := make(Chan[address])
	go myPrint(addrch)
	for _, a := range addrList {
		addrch <- a
	}
}
