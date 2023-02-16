package main

import "fmt"

func add(num int) {
	num += 1
}

func add_val(num *int) {
	*num += 1
}

func main() {
	// map
	m1 := make(map[string]int)
	m2 := map[string]string{
		"Sam": "Male",
		"Alice": "female",
	}
	m1["tom"] = 18
	fmt.Println(m1)
	fmt.Println(m2["Sam"], m2["Alice"])

	// pointer 
	str := "Golang"
	var p *string = &str
	fmt.Println(*p)
	*p = "hello"
	fmt.Println(*p, str)

	// Go使用的是值传递
	num := 100
	add(num)
	fmt.Printf("num = %d\n", num)

	add_val(&num)
	fmt.Printf("num = %d\n", num)

	return
}