package _case

import (
	"fmt"
	"io"
	"log"
	"os"
)

// DeferCase1 defer 关键词用来声明一个延迟调用函数
// 该函数可以是匿名函数也可以是具名函数
// 延迟调用函数在return 之后，main函数拿到返回值之前进行调用
// defer延迟函数的执行顺序为后进先出
func DeferCase1() {
	fmt.Println("开始执行DeferCase1")
	defer func() {
		fmt.Println("调用了 匿名函数1")
	}()
	defer f1()
	defer func() {
		fmt.Println("调用了 匿名函数3")
	}()

	fmt.Println("DeferCase1执行结束")
}

// DeferCase2 参数预计算
func DeferCase2() {
	i := 1
	// 传参，go中函数是一个list，有参函数引用的是按照顺序的局部变量的值
	defer func(j int) {
		fmt.Println("defer j : ", j)
	}(i + 1) //加（）代表声明完函数立即执行，无需再其他处理即可使用

	// 闭包，访问函数内局部变量，即函数return之后的局部变量的值
	defer func() {
		fmt.Println("defer i : ", i)
	}()
	i = 99
	fmt.Println("i:", i)
}

// DeferCase3 返回值
// defer 函数执行在return之后
func DeferCase3() {
	i, j := f2()
	fmt.Printf("i : %d, j : %d, g : %d\n", i, *j, g)
}

func f1() {
	fmt.Println("调用了具名函数f1")
}

var g = 100

func f2() (int, *int) {
	defer func() {
		g = 200
	}()
	fmt.Println("f2 g:", g)
	return g, &g
}

func ExceptionCase() {
	defer func() {
		// 恢复协程，捕获异常
		err := recover()
		// 异常处理
		if err != nil {
			fmt.Println("异常处理 defer recover:", err)
		}
	}()

	fmt.Println("开始执行ExceptionCase")
	panic("ExceptionCase抛出异常")
	fmt.Println("ExceptionCase结束执行")
}

func FileReadCase() {
	file, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	// 通过 defer 调用资源释放的方法
	defer func() {
		file.Close()
		fmt.Println("释放文件资源")
	}()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(buf[:n])
	}
}
