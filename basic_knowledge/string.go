package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "Golang"
	var str2 string = "Go语言"

	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Println(str1[2], rune(str1[2]))
	fmt.Printf("%d %c\n", str2[2], str2[2])
	fmt.Println("len(str2)：", len(str2))
	return
}