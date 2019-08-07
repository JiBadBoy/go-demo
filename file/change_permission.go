package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// 使用Linux风格改变文件权限
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}
	// 改变文件所有者 windows下可能出问题
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}
	// 改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}
