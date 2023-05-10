package _case

import (
	"fmt"
	"os"
)

func FmtCase() {
	fmt.Println("今天天气很好")
	fmt.Printf("%s天气很好\n", "今天")
	str := fmt.Sprintf("%s天气很好", "今天、明天")
	fmt.Fprint(os.Stderr, str)
}

func FmtCase1() {
	type simple struct {
		value int
	}
	a := simple{
		value: 10,
	}

	// 通用占位符
	fmt.Printf("默认格式的值：%v \n", a)
	fmt.Printf("包含字段名的值：%+v \n", a)
	fmt.Printf("go语法表示的值：%#v \n", a)
	fmt.Printf("go语法表示的类型：%T \n", a)
	fmt.Printf("输出字面上的百分号：10%% \n")

	// 整数占位符
	v1 := 10
	v2 := 20170 // "今" 字码点值
	fmt.Printf("二进制：%b\n", v1)
	fmt.Printf("Unicode 码点值转字符:%c\n", v2)
	fmt.Printf("十进制：%d\n", v1)
	fmt.Printf("八进制：%o\n", v1)
	fmt.Printf("以0o为前缀的八进制：%O \n", v1)
	fmt.Printf("用单引号将字符的值包起来：%q \n", v2)
	fmt.Printf("十六进制：%x \n", v1)
	fmt.Printf("十六进制大写：%X \n", v1)
	fmt.Printf("Unicode 格式：%U \n", v2)

	// 宽度设置
	fmt.Printf("%v 的二进制 %b; go语言语法表示二进制为%#b; 指定二进制宽度为8， 不足8位补0：%08b \n", v1, v1, v1, v1)
	fmt.Printf("%v 的十六进制 %x; go语言语法表示指定十六进制宽度为8， 不足8位补0：%#08x \n", v1, v1, v1)
	fmt.Printf("%v 的字符表示为 %c; 指定宽度为5， 不足5位用空格表示：%5c \n", v2, v2, v2)

	// 浮点数占位符
	var f1 = 123.456
	var f2 = 1234567890.78999
	fmt.Printf("指数为二的幂的无小数科学计算法: %b \n", f1)
	fmt.Printf("科学计算法: %e \n", f1)
	fmt.Printf("大写的科学计算法: %E \n", f1)
	fmt.Printf("有小数点而无指数，即常规的浮点数格式，默认宽度和精度: %f \n", f1)
	fmt.Printf("宽度为9，精度默认为：%9f \n", f1)
	fmt.Printf("默认宽度，精度保留两位小数：%.2f \n", f1)
	fmt.Printf("宽度为9，精度保留两位小数：%9.2f \n", f1)
	fmt.Printf("宽度为9，精度保留0位小数：%9.f \n", f1)
	fmt.Printf("根据情况自动选择%%e 或者 %%f 来输出，以产生更紧凑的输出（末位无0）：%g %g \n", f1, f2)
	fmt.Printf("根据情况自动选择%%E 或者 %%f 来输出，以产生更紧凑的输出（末位无0）：%G %G \n", f1, f2)
	fmt.Printf("以十六进制方式表示：%x \n", f1)
	fmt.Printf("以十六进制大写方式表示：%X \n", f1)

	// 字符串占位符
	var str = "今天是个好日子"
	fmt.Printf("描述一下今天：%s \n", str)
	// 用双引号将字符串包裹
	fmt.Printf("描述一下今天：%q \n", str)
	// 用十六进制表示
	fmt.Printf("%x \n", str)
	// 以空格作为两数之间分隔符，并用大写十六进制表示
	fmt.Printf("% X \n", str)

	// 指针占位符
	var str1 = "今天是个好日子"
	bytes := []byte(str1)
	// 切片第0个元素地址
	fmt.Printf("%p \n", bytes)
	mp := make(map[string]int, 0)
	fmt.Printf("%p \n", mp)
	var p *map[string]int = new(map[string]int)
	fmt.Printf("%p \n", p)
}
