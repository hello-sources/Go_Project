package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	f1()
	f2()
	gin.New()
}

func f1() {
	fmt.Println("调用方法1")
}
