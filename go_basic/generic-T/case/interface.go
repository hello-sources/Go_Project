package _case

import "fmt"

// ToString 基本接口，可用于变量定义
type ToString interface {
	String() string
}

// 基本接口定义一个变量
var ts ToString

func (u user) String() string {
	return fmt.Sprintf("ID : %d, Name : %s, Age : %d", u.ID, u.Name, u.Age)
}

func (addr address) String() string {
	return fmt.Sprintf("ID : %d, Province : %s, City : %s", addr.ID, addr.Province, addr.City)
}

// GetKey 泛型接口
type GetKey[T comparable] interface {
	any
	Get() T
}

func (u user) Get() int64 {
	return u.ID
}
func (addr address) Get() int {
	return addr.ID
}

// 列表转集合
func listToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(MapT[k, T], len(list))
	for _, data := range list {
		mp[data.Get()] = data
	}
	return mp
}

func InterfaceCase() {
	userList := []GetKey[int64]{
		user{ID: 1, Name: "nick", Age: 18},
		user{ID: 2, Name: "tom", Age: 19},
	}

	addrList := []GetKey[int]{
		address{ID: 1, Province: "安徽", City: "合肥"},
		address{ID: 2, Province: "江苏", City: "南京"},
	}

	userMp := listToMap[int64, GetKey[int64]](userList)
	fmt.Println(userMp)
	addrMp := listToMap[int, GetKey[int]](addrList)
	fmt.Println(addrMp)

}
