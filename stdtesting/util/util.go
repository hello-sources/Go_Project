package util

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

const defaultNum = 100

func ParseSize(size string) (int64, string) {
	//time.Sleep(time.Nanosecond * 500)
	re, _ := regexp.Compile("[0-9]+")
	uint := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, uint, "", 1), 10, 64)
	uint = strings.ToUpper(uint)
	var byteNum int64 = 0
	switch uint {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
	}

	if num == 0 {
		log.Println("ParseSize 仅支持B、KB、MB、GB、TB、PB")
		num = defaultNum
		byteNum = num * MB
		uint = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + uint
	return byteNum, sizeStr
}
