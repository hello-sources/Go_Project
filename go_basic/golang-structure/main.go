// 声明包
package main

// 引入包
import (
	"fmt"
	_case "goBasic/generic-T/case"
	operator "goBasic/golang-structure/case"
	_case1 "goBasic/type-var/case"
)

// 初始化函数，golang每个包的引用会优先调用该函数
func init() {
	fmt.Println("调用 golang-structure init")
}

// 函数(main函数为整个程序入口)
func main() {
	var i = 10
	// 语句或表达式
	fmt.Println("调用 golang-structure main")
	fmt.Println(fmt.Sprintf("打印参数i : %d", i))
	// 包名可以与包引用目录不一致（包名和文件夹名称可以不一致）
	// golang公有成员和私有成员通过成员表示的首字母来确认
	// 首字母大写表示公有成员，首字母小写表示私有成员
	_case.Simplecase()
	_case1.ConstAndEnumCase()

	operator.ArithmeticCase()
	operator.RelationCase()
	operator.LogicCase()
	operator.BitCase()
	operator.AssignmentCase()
}

/*
这是多行注释
这是多行注释
这是多行注释
*/
