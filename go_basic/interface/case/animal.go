package _case

import "fmt"

// Animal
// 声明 Animal 接口
// 定义 Animal 行为
type Animal interface {
	Each()
	Drink()
	Sleep()
	Run()
}

type animal struct {
}

func (a animal) Each() {
	fmt.Println(" Animal Each 接口默认实现")
}

func (a animal) Drink() {
	fmt.Println(" Animal Drink 接口默认实现")

}

func (a animal) Sleep() {
	fmt.Println(" Animal Sleep 接口默认实现")

}

func (a animal) Run() {
	fmt.Println(" Animal Run 接口默认实现")
}
