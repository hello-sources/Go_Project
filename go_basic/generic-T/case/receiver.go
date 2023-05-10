package _case

import "fmt"

// MyStruct 泛型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// GetData 泛型receiver
func (myStruct MyStruct[T]) GetData() T {
	return myStruct.Data
}

func ReceiverCase() {
	data := 18
	myStruct := MyStruct[*int]{
		Name: "nick",
		Data: &data,
	}
	data1 := myStruct.GetData()
	fmt.Println(*data1)

	str := "abcdefg"
	myStruct1 := MyStruct[*string]{
		Name: "nick",
		Data: &str,
	}
	str1 := myStruct1.GetData()
	fmt.Println(*str1)
}
