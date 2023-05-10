package _case

import (
	"fmt"
	"regexp"
)

func RegexpCase() {
	// 构建一个正则表达式的对象
	reg := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	// 判断给定的字符串是否符合正则
	fmt.Println(reg.MatchString("abcd[1234]"))
	// 从给定的字符串查找符合条件的字符串
	bytes := reg.FindAll([]byte("1234efg[456]dab"), -1)
	if bytes != nil {
		fmt.Println(string(bytes[0]))
	} else {
		fmt.Println("没找到合适的字符串")
	}
}
