package main

import (
	"fmt"
)

func main() {
	slice1 := make([]float32, 0)
	slice2 := make([]int, 3, 5)
	fmt.Printf("slice1 = %s\n", slice1)
	fmt.Println(len(slice2), cap(slice2))

	slice2 = append(slice2, 2, 3, 4, 5)
	fmt.Println(len(slice2), cap(slice2))

	sub1 := slice2[3:5]
	sub2 := slice2[2:]
	sub3 := slice2[:5]
	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
	return
}