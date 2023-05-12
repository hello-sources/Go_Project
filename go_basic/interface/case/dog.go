package _case

import "fmt"

type Dog struct {
	animal
}

func NewDog() Animal {
	//return Dog{}
	return &Dog{}
}

func (dog Dog) Each() {
	fmt.Println("狗吃肉包子")
}

func (dog Dog) Drink() {
	fmt.Println("狗喝白开水")
}

func (dog Dog) Sleep() {
	fmt.Println("狗睡在窝里")
}

func (dog Dog) Run() {
	fmt.Println("狗撒欢的跑")
}
