package main

import "fmt"

type Gender int8

const (
	MALE Gender = 1
	FAMALE Gender = 2
)

func main() {
	young, old := 18, 20
	if young > 16 {
		fmt.Println("A young man")
	} else {
		fmt.Println("too little")
	}
	if old > 18 {
		fmt.Println("Old enougth")
	}

	gender := MALE
	switch gender {
	case MALE:
		fmt.Println("你是男人")
		fallthrough
	case FAMALE:
		fmt.Println("你是女人")
	default:
		fmt.Println("你很危险，小老弟")	
	}

	nums := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	
	for key, value := range m2 {
		fmt.Println(key, value)
	}

	return
}
