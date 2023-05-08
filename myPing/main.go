package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	size    int
	count   int
	icmp    *ICMP = &ICMP{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Id:          1,
		SequenceNum: 1,
	}
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Id          uint16
	SequenceNum uint16
}

func main() {
	var a uint16 = 33
	var b uint16 = 32
	fmt.Println(a + b)

	getCommandArgs()
	fmt.Println(timeout, size, count)

	desIP := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", desIP, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	fmt.Println(buffer)

	data := make([]byte, size)
	fmt.Println(data)
}

func getCommandArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长，单位毫秒")
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小，单位字节")
	flag.IntVar(&count, "n", 4, "发送请求数")
	flag.Parse()
}

//42 分钟 https://www.bilibili.com/video/BV1e24y1T7P4?p=67&vd_source=3dd696fe29b7bd6acd4fca4daf216a36
