package _case

import (
	"log"
	"os"
	"strings"
)

// 源文件目录
const sourceDir = "D:/Software/Go_Project/Go_Project/go_basic/file-rw/source-file/"

// 目标文件目录
const destDir = "D:/Software/Go_Project/Go_Project/go_basic/file-rw/dest-file/"

func getFiles(dir string) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	list := make([]string, 0)
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		fullName := strings.Trim(dir, "/") + "/" + f.Name()
		list = append(list, fullName)
	}
	return list
}
