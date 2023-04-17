package main

import "fmt"

func main() {
	println("test")
	println(`
		test 
		"换行"
	`)
	println(len("test"))
	println(len("你"))
}