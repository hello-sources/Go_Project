package _case

import "fmt"

const (
	FEMALE Gender = iota
	MALE
	THIRD
	UNKNOWN
)

type Gender uint8

type Person interface {
	Run()
	Sleep()
}

type Teacher struct {
	Name   string
	Age    uint8
	Gender Gender
}

func (Teacher) Run() {
	fmt.Println("老师在公园 run")
}

func (Teacher) Sleep() {
	fmt.Println("老师在家 sleep")
}

func PersonCase(person Person) {
	defer func() {
		fmt.Println("person defer")
	}()
	person.Run()
	person.Sleep()
}

func PersonCase1(person interface{}) {
	p1, ok := person.(Person)
	if ok {
		p1.Run()
	} else {
		fmt.Println("类型不能识别")
	}
}

type Student struct {
}

func (Student) Run() {
	fmt.Println("学生在操场 run")
}

func (u *Student) Sleep() {
	fmt.Println("学生在寝室 sleep")
}

func TryCatchCase() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	PanicCase()
}

func PanicCase() {
	panic("程序出现异常")
}

func DeferCase() {
	i := 1
	// 传参，传参只采用当前程序下的参数值
	defer func(j int) {
		fmt.Println("defer j:", j)
	}(i + 1)

	// 闭包，使用最后的外部变量的结果
	defer func() {
		i++
		fmt.Println("defer i :", i)
	}()
	i = 99
	return
}

func DeferCase1() {
	i, i1 := f1()
	fmt.Println(i, *i1, j)
}

var j int = 1

func f1() (int, *int) {
	defer func() {
		j = 100
	}()
	fmt.Println("f1 j : ", j)
	return j, &j
}
