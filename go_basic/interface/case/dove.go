package _case

import "fmt"

type Dove struct {
}

func NewDove() Animal {
	return &Dove{}
}

func (d *Dove) Each() {
	fmt.Println("鸽子吃虫子")
}

func (d *Dove) Drink() {
	fmt.Println("鸽子也喝白开水")
}

func (d *Dove) Sleep() {
	fmt.Println("鸽子睡觉")
}

func (d *Dove) Run() {
	fmt.Println("鸽子飞")
}
